package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetCotizacion(monId *string, fchCotiz *string) (*fe.FEParamGetCotizacionResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetCotizacion{Auth: *authData,
		MonId:    monId,
		FchCotiz: fchCotiz,
	}
	// Request
	return request[fe.FEParamGetCotizacion, fe.FEParamGetCotizacionResponse](structSend)
}
