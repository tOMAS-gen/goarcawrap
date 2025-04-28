package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetTiposPaises() (*fe.FEParamGetTiposPaisesResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetTiposPaises{Auth: *authData}
	// Request
	return request[fe.FEParamGetTiposPaises, fe.FEParamGetTiposPaisesResponse](structSend)
}
