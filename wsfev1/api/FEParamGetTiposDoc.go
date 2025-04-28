package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetTiposDoc() (*fe.FEParamGetTiposDocResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetTiposDoc{Auth: *authData}
	// Request
	return request[fe.FEParamGetTiposDoc, fe.FEParamGetTiposDocResponse](structSend)
}
