package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetTiposOpcional() (*fe.FEParamGetTiposOpcionalResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetTiposOpcional{Auth: *authData}
	// Request
	return request[fe.FEParamGetTiposOpcional, fe.FEParamGetTiposOpcionalResponse](structSend)
}
