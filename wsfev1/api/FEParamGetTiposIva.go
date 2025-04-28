package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetTiposIva() (*fe.FEParamGetTiposIvaResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetTiposIva{Auth: *authData}
	// Request
	return request[fe.FEParamGetTiposIva, fe.FEParamGetTiposIvaResponse](structSend)
}
