package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecaea"
)

func FECAEASinMovimientoConsultar(caea *string, ptoVta int) (*fecaea.FECAEASinMovimientoConsultarResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fecaea.FECAEASinMovimientoConsultar{Auth: *authData,
		CAEA:   caea,
		PtoVta: ptoVta,
	}
	// Request
	return request[fecaea.FECAEASinMovimientoConsultar, fecaea.FECAEASinMovimientoConsultarResponse](structSend)
}
