package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetPtosVenta() (*fe.FEParamGetPtosVentaResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetPtosVenta{Auth: *authData}
	// Request
	return request[fe.FEParamGetPtosVenta, fe.FEParamGetPtosVentaResponse](structSend)
}
