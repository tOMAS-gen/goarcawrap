package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetTiposCbte() (*fe.FEParamGetTiposCbteResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetTiposCbte{Auth: *authData}
	// Request
	return request[fe.FEParamGetTiposCbte, fe.FEParamGetTiposCbteResponse](structSend)
}
