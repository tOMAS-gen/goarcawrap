package wsaa

import (
	"fmt"

	"github.com/tOMAS-gen/goarcawrap/model"
)

func Authenticate() (*model.WSAA, error) {
	// --- Llamada al login ---
	loginTicketResponse, err := Client()
	if err != nil {
		return nil, fmt.Errorf("Error durante el login en WSAA: %v", err)
	}
	// --- Procesar respuesta ---
	wsaa, err := parce(loginTicketResponse)
	if err != nil {
		return nil, fmt.Errorf("Error al procesar la respuesta del WSAA: %v", err)
	}
	// --- Guardar respuesta ---
	err = Save(wsaa)
	if err != nil {
		return nil, fmt.Errorf("Error al guardar la respuesta del WSAA: %v", err)
	}
	// Devuelve el objeto WSAA
	return wsaa, nil
}
