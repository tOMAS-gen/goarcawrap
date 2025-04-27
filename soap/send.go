package soap

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SendSoapRequest envía una solicitud SOAP al endpoint de WSFEV1.
// Recibe el cuerpo XML de la solicitud como un string.
// Devuelve el cuerpo de la respuesta como un string y un error si ocurre alguno.
func SendSoapRequest(xmlRequestBody string, urlApi string) (string, error) {
	// 1. Crear un reader para el cuerpo de la solicitud XML
	//    strings.NewReader es eficiente para leer desde un string
	requestBodyReader := strings.NewReader(xmlRequestBody)

	// 2. Crear la petición HTTP POST
	//    http.NewRequest te da más control sobre la petición, incluyendo las cabeceras.
	req, err := http.NewRequest("POST", urlApi, requestBodyReader)
	if err != nil {
		return "", fmt.Errorf("error al crear la petición HTTP: %w", err)
	}

	// 3. Establecer la cabecera Content-Type requerida por SOAP 1.2
	//    Es crucial para que el servidor entienda el formato del cuerpo.
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	// Podrías necesitar añadir otras cabeceras aquí si el servicio las requiere,
	// como SOAPAction para SOAP 1.1 (aunque WSFEV1 usualmente usa SOAP 1.2 sin SOAPAction explícita en la cabecera).

	// 4. Crear un cliente HTTP
	//    Puedes configurar timeouts y otras opciones en el cliente si es necesario.
	client := &http.Client{}

	// 5. Enviar la petición
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error al enviar la petición HTTP: %w", err)
	}
	// Asegurarse de cerrar el cuerpo de la respuesta al final de la función
	defer resp.Body.Close()

	// 6. Leer el cuerpo de la respuesta
	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer el cuerpo de la respuesta: %w", err)
	}
	responseBodyString := string(responseBodyBytes)

	// 7. Verificar el código de estado HTTP
	//    Un código 2xx indica éxito, otros pueden indicar errores del lado del servidor.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return responseBodyString, fmt.Errorf("error en la respuesta del servidor (Código de estado: %d) info: %v", resp.StatusCode, responseBodyString)
	}

	// 8. Devolver el cuerpo de la respuesta como string
	return responseBodyString, nil
}
