package wsaa

import (
	"fmt"

	"github.com/tOMAS-gen/goarcawrap/model"
)

var wsaa = make(map[string]model.WSAA)

func get(serviceID string) (*model.WSAA, error) {
	// Variable
	var wsaaSend *model.WSAA
	// Verificar si esta cargado
	wsaaMap, ok := wsaa[serviceID]
	if ok {
		// Remplazar
		wsaaSend = &wsaaMap
		// Si esta cargado, verificar si está vigente
		if !checkExpirationTimeWSAA(&wsaaMap) {
			// Si no está vigente, renovar
			wsaaAux, err := authenticate(serviceID)
			if err != nil {
				return nil, err
			}
			// devolver
			wsaaSend = wsaaAux
		}
	} else {
		// Si no esta cargado, cargar
		wsaaAux, err := readAndVerify(serviceID)
		if err != nil {
			return nil, err
		}
		// Guardar
		wsaaSend = wsaaAux
		wsaa[serviceID] = *wsaaSend
	}
	return wsaaSend, nil
}

func readAndVerify(serviceID string) (*model.WSAA, error) {
	// Verificar si existe
	wsaaAux, err := read(serviceID)
	if err != nil {
		// Si no existe, generar
		return authenticate(serviceID)
	}
	// Comprovar si está vigente
	if !checkExpirationTimeWSAA(wsaaAux) {
		// Si no está vigente, renovar
		return authenticate(serviceID)
	}
	// Devolver WSAA
	return wsaaAux, nil
}

func PrintData() {
	for key, value := range wsaa {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
