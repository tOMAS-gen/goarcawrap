package wsaa

import "github.com/tOMAS-gen/goarcawrap/model"

var wsaa *model.WSAA

func Get() (*model.WSAA, error) {
	// Verificar si esta cargado
	if wsaa == nil {
		// Si no esta cargado, cargar
		wsaaAux, err := read()
		if err != nil {
			return nil, err
		}
		// devolver
		wsaa = wsaaAux
	} else {
		// Si esta cargado, verificar si está vigente
		if !CheckExpirationTimeWSAA(wsaa) {
			// Si no está vigente, renovar
			wsaaAux, err := Authenticate()
			if err != nil {
				return nil, err
			}
			// devolver
			wsaa = wsaaAux
		}
	}
	return wsaa, nil
}

func read() (*model.WSAA, error) {
	// Verificar si existe
	wsaaAux, err := Read()
	if err != nil {
		// Si no existe, generar
		return Authenticate()
	}
	// Comprovar si está vigente
	if !CheckExpirationTimeWSAA(wsaaAux) {
		// Si no está vigente, renovar
		return Authenticate()
	}
	// Devolver WSAA
	return wsaaAux, nil
}
