package wsaa

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/tOMAS-gen/goarcawrap/model"
	"go.mozilla.org/pkcs7"
)

const (
	// Formato de fecha/hora requerido por AFIP (similar a ISO 8601)
	afipTimeFormat = "2006-01-02T15:04:05-07:00"
	// Formato para UniqueID (similar a PowerShell "yyMMddHHMM")
	uniqueIDFormat = "0601021504"
	// Duración del ticket solicitado (ej. +/- 10 minutos desde ahora)
	ticketDuration = 10 * time.Minute
	// Namespace del servicio WSAA
	wsaaWsdlURL = model.URL_wsaa
	// Namespace del servicio WSFE
	// serviceID = "wsfe"
)

// Estructura para el Login Ticket Request (TRA) XML
type LoginTicketRequest struct {
	XMLName xml.Name `xml:"loginTicketRequest"`
	Version string   `xml:"version,attr,omitempty"` // Opcional, AFIP lo ignora pero puede estar
	Header  Header   `xml:"header"`
	Service string   `xml:"service"`
}

type Header struct {
	UniqueID       string    `xml:"uniqueId"`
	GenerationTime time.Time `xml:"generationTime"`
	ExpirationTime time.Time `xml:"expirationTime"`
}

// Estructuras para el sobre SOAP de la solicitud
type SoapEnvelopeRequest struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SoapBodyRequest `xml:"Body"`
}

type SoapBodyRequest struct {
	XMLName  xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	LoginCms LoginCmsRequest `xml:"http://wsaa.view.sua.dvadac.desein.afip.gov loginCms"` // Namespace del servicio WSAA
}

type LoginCmsRequest struct {
	In0 string `xml:"in0"` // Contiene el CMS base64
}

// Estructuras para el sobre SOAP de la respuesta
type SoapEnvelopeResponse struct {
	XMLName xml.Name         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SoapBodyResponse `xml:"Body"`
}

type SoapBodyResponse struct {
	XMLName        xml.Name         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	LoginCmsReturn LoginCmsResponse `xml:"loginCmsResponse"`
}

type LoginCmsResponse struct {
	XMLName xml.Name `xml:"loginCmsResponse"`
	Result  string   `xml:"loginCmsReturn"` // Contiene el Login Ticket Response XML o un error SOAP Fault
}

func Client(serviceID string) (*string, error) {
	// Definir flags de línea de comandos (parámetros)
	certPath := filepath.Join(model.DataDir, model.CRTfileName)
	keyPath := filepath.Join(model.DataDir, model.PrivateKeyFileName)

	// --- Cliente WSAA AFIP en Go ---

	// --- PASO 0: Cargar Certificado y Clave Privada ---
	certBytes, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo de certificado: %v", err)
	}

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo de clave privada: %v", err)
	}

	// Parsear el certificado PEM
	pemBlock, _ := pem.Decode(certBytes)
	if pemBlock == nil {
		return nil, fmt.Errorf("Error al decodificar el bloque PEM del certificado")
	}
	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Error al parsear el certificado: %v", err)
	}

	// Parsear la clave privada PEM
	pemBlock, _ = pem.Decode(keyBytes)
	if pemBlock == nil {
		return nil, fmt.Errorf("Error al decodificar el bloque PEM de la clave privada")
	}

	var privKey *rsa.PrivateKey
	// Intentar parsear como PKCS#1
	if key, errPkcs1 := x509.ParsePKCS1PrivateKey(pemBlock.Bytes); errPkcs1 == nil {
		privKey = key
	} else {
		// Si falla, intentar parsear como PKCS#8
		if key, errPkcs8 := x509.ParsePKCS8PrivateKey(pemBlock.Bytes); errPkcs8 == nil {
			var ok bool
			privKey, ok = key.(*rsa.PrivateKey)
			if !ok {
				return nil, fmt.Errorf("La clave privada PKCS#8 no es de tipo RSA")
			}
		} else {
			return nil, fmt.Errorf("Error al parsear la clave privada (no es PKCS#1 ni PKCS#8 válida): %v, %v", errPkcs1, errPkcs8)
		}
	}

	// --- PASO 1: ARMAR EL XML DEL TICKET DE ACCESO (TRA) ---
	now := time.Now()
	genTime := now.Add(-ticketDuration) // Tiempo de generación un poco antes
	expTime := now.Add(ticketDuration)  // Tiempo de expiración un poco después

	// Crear la estructura del TRA
	tra := LoginTicketRequest{
		Header: Header{
			UniqueID: fmt.Sprintf("%d", now.Unix()), // Usar timestamp Unix como UniqueID simple
			// UniqueID:       now.Format(uniqueIDFormat), // Alternativa formato YYMMDDHHMM
			GenerationTime: genTime.UTC(), // Usar UTC
			ExpirationTime: expTime.UTC(), // Usar UTC
		},
		Service: serviceID,
	}

	// Generar el XML del TRA en memoria
	xmlTRA, err := xml.MarshalIndent(tra, "", "  ") // MarshalIndent para debug, quitar Indent para producción
	if err != nil {
		return nil, fmt.Errorf("Error al generar el XML del TRA: %v", err)
	}
	// Reemplazar los formatos de fecha/hora generados por Go con los requeridos por AFIP
	// xml.Marshal usa RFC3339Nano, AFIP necesita un formato específico sin nano y con timezone.
	xmlTRAStr := string(xmlTRA)
	xmlTRAStr = string(bytes.Replace([]byte(xmlTRAStr), []byte(tra.Header.GenerationTime.Format(time.RFC3339Nano)), []byte(tra.Header.GenerationTime.Format(afipTimeFormat)), 1))
	xmlTRAStr = string(bytes.Replace([]byte(xmlTRAStr), []byte(tra.Header.ExpirationTime.Format(time.RFC3339Nano)), []byte(tra.Header.ExpirationTime.Format(afipTimeFormat)), 1))
	xmlTRA = []byte(xmlTRAStr)

	// --- PASO 2: FIRMAR CMS ---

	// Convertir la clave privada genérica a *rsa.PrivateKey si es necesario
	// (ya lo hicimos al cargarla, pero por si acaso)
	// Since we already validated privKey is an RSA key when loading it,
	// we can directly use it without another type assertion
	rsaPrivKey := privKey
	// Crear el objeto SignedData de PKCS#7
	signedData, err := pkcs7.NewSignedData(xmlTRA)
	if err != nil {
		return nil, fmt.Errorf("Error al inicializar SignedData de PKCS#7: %v", err)
	}

	// Añadir información del firmante (certificado y clave)
	// Usar rand.Reader para la aleatoriedad necesaria en la firma
	if err = signedData.AddSigner(cert, rsaPrivKey, pkcs7.SignerInfoConfig{}); err != nil {
		return nil, fmt.Errorf("Error al añadir el firmante al CMS: %v", err)
	}

	// Finalizar la firma (modo detached=false implícito por defecto)
	// Obtener los bytes del CMS en formato DER
	derCMS, err := signedData.Finish()
	if err != nil {
		return nil, fmt.Errorf("Error al finalizar la firma CMS: %v", err)
	}

	// --- PASO 3: ENCODEAR EL CMS EN BASE 64 ---
	base64CMS := base64.StdEncoding.EncodeToString(derCMS)
	// log.Println("CMS Base64:", base64CMS) // Descomentar para depurar

	// --- PASO 4: INVOCAR AL WSAA ---
	// Extraer la URL del endpoint del servicio desde la URL del WSDL
	// Usualmente es la URL base sin "?WSDL"
	endpointURL := wsaaWsdlURL
	if wsdlQueryParam := "?WSDL"; len(endpointURL) > len(wsdlQueryParam) && endpointURL[len(endpointURL)-len(wsdlQueryParam):] == wsdlQueryParam {
		endpointURL = endpointURL[:len(endpointURL)-len(wsdlQueryParam)]
	}

	// Crear el cuerpo de la solicitud SOAP
	soapReq := SoapEnvelopeRequest{
		Body: SoapBodyRequest{
			LoginCms: LoginCmsRequest{
				In0: base64CMS,
			},
		},
	}

	// Generar el XML de la solicitud SOAP
	soapReqBytes, err := xml.MarshalIndent(soapReq, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Error al generar el XML de la solicitud SOAP: %v", err)
	}

	// Crear la solicitud HTTP POST
	httpReq, err := http.NewRequest("POST", endpointURL, bytes.NewBuffer(soapReqBytes))
	if err != nil {
		return nil, fmt.Errorf("Error al crear la solicitud HTTP: %v", err)
	}

	// Establecer las cabeceras HTTP necesarias para SOAP
	httpReq.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpReq.Header.Set("SOAPAction", `""`) // WSAA LoginCms generalmente no requiere una SOAPAction específica

	// Ejecutar la solicitud HTTP
	httpClient := &http.Client{Timeout: 30 * time.Second} // Timeout razonable
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("Error al enviar la solicitud HTTP al WSAA: %v", err)
	}
	defer httpResp.Body.Close() // Asegurarse de cerrar el cuerpo de la respuesta

	// Leer el cuerpo de la respuesta HTTP
	respBodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error al leer cuerpo de la respuesta del WSAA: %v", err)
	}

	// Verificar si hubo un error HTTP (ej. 4xx, 5xx)
	if httpResp.StatusCode >= 400 {
		// Imprimir el cuerpo de la respuesta para diagnóstico
		return nil, fmt.Errorf("Error HTTP %d recibido del servidor.\n Cuerpo de la respuesta cruda:\n %v", httpResp.StatusCode, string(respBodyBytes))
	}

	// Parsear la respuesta SOAP XML
	var soapResp SoapEnvelopeResponse
	err = xml.Unmarshal(respBodyBytes, &soapResp)
	if err != nil {
		// A veces, si AFIP devuelve un error SOAP Fault, el Unmarshal directo a nuestra estructura puede fallar.
		// Imprimimos el cuerpo crudo para diagnóstico.
		return nil, fmt.Errorf("Error al parsear la respuesta SOAP XML: %s\n Cuerpo de la respuesta cruda:\n %v", err, string(respBodyBytes))

	}
	// Imprimir el resultado
	println(soapResp.Body.LoginCmsReturn.Result)
	// --- PASO 5: DEVOLVER RESULTADO  ---
	return &soapResp.Body.LoginCmsReturn.Result, nil
}
