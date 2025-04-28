package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecaea"
)

func FECAEASinMovimientoInformar(caea *string, ptoVta int) (*fecaea.FECAEASinMovimientoInformarResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fecaea.FECAEASinMovimientoInformar{Auth: *authData,
		CAEA:   caea,
		PtoVta: ptoVta,
	}
	// Request
	return request[fecaea.FECAEASinMovimientoInformar, fecaea.FECAEASinMovimientoInformarResponse](structSend)
}
