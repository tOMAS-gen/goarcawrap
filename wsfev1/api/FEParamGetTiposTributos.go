package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetTiposTributos() (*fe.FEParamGetTiposTributosResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetTiposTributos{Auth: *authData}
	// Request
	return request[fe.FEParamGetTiposTributos, fe.FEParamGetTiposTributosResponse](structSend)
}
