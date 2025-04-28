package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecaea"
)

func FECAEASolicitar(periodo int, orden int16) (*fecaea.FECAEASolicitarResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fecaea.FECAEASolicitar{Auth: *authData,
		Periodo: periodo,
		Orden: orden,
	}
	// Request
	return request[fecaea.FECAEASolicitar, fecaea.FECAEASolicitarResponse](structSend)
}
