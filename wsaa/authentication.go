package wsaa

import (
	"fmt"
)

func Authenticate() (*string, error) {
	// --- Llamada al login ---
	loginTicketResponse, err := Client()
	if err != nil {
		return nil, fmt.Errorf("Error durante el login en WSAA: %v", err)
	}
	// --- Procesar la respuesta ---
	return loginTicketResponse, nil
}