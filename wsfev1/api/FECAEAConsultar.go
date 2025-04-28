package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecaea"
)

func FECAEAConsultar(periodo int, orden int16) (*fecaea.FECAEAConsultarResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fecaea.FECAEAConsultar{Auth: *authData,
		Periodo: periodo,
		Orden:   orden,
	}
	// Request
	return request[fecaea.FECAEAConsultar, fecaea.FECAEAConsultarResponse](structSend)
}
