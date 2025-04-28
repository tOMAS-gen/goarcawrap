package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetActividades() (*fe.FEParamGetActividadesResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetActividades{Auth: *authData}
	// Request
	return request[fe.FEParamGetActividades, fe.FEParamGetActividadesResponse](structSend)
}