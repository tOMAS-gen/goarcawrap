package model

import "time"

const (
	DataDir            = "ARCA"               // Directorio base para guardar datos de clientes
	PrivateKeyFileName = "MiClavePrivada.key" // Nombre para el archivo de clave privada
	CSRfileName        = "MiPedido.csr"       // Nombre para el archivo CSR
	CRTfileName        = "MiCertificado.crt"  // Nombre para el archivo CSR
	KeyBits            = 2048                 // Tamaño de la clave RSA (mínimo 2048 para AFIP/ARCA)
	ClientFileName     = "client.json"        // Nombre del archivo JSON
	WSAAfileName       = "wsaa.json"          // Nombre del archivo JSON
	// // production
	// URL_wsaa             = "https://wsaa.afip.gov.ar/ws/services/LoginCms"
	// URL_wsfe1            = "https://servicios1.afip.gov.ar/wsfev1/service.asmx"
	// URL_ws_sr_padron_a13 = "https://aws.afip.gov.ar/sr-padron/webservices/personaServiceA13"
	// URL_ws_sr_constancia_inscripcion = "https://aws.afip.gov.ar/sr-padron/webservices/personaServiceA5"
	// testing/homologación
	URL_wsaa = "https://wsaahomo.afip.gov.ar/ws/services/LoginCms"
	URL_wsfe1 = "https://wswhomo.afip.gov.ar/wsfev1/service.asmx"
	URL_ws_sr_padron_a13 = "https://awshomo.afip.gov.ar/sr-padron/webservices/ersonaServiceA13"
	URL_ws_sr_constancia_inscripcion = "https://awshomo.afip.gov.ar/sr-padron/webservices/personaServiceA5"
)

type Client struct {
	CUIT             int64  `json:"cuit"`
	CommonName       string `json:"commonName"`
	OrganizationName string `json:"organizationName"`
}

type WSAA struct {
	Token          string    `json:"token"`
	Sign           string    `json:"sign"`
	ExpirationTime time.Time `json:"expireAt"`
}
