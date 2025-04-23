package wsaa

import "github.com/tOMAS-gen/goarcawrap/model"

func Get() (*model.WSAA, error) {
	// Verificar si existe
	wsaa, err := Read()
	if err!= nil {
		// Si no existe, generar
		return Authenticate()
	}
	// Comprovar si está vigente
	if !CheckExpirationTimeWSAA(wsaa) {
		// Si no está vigente, renovar
		return Authenticate()
	}
	// Devolver WSAA
	return wsaa, nil
}

