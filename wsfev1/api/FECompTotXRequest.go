package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/auth"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FECompTotXRequest() (*fe.FECompTotXRequestResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FECompTotXRequest{Auth: auth.FEAuthRequest{
		Token: authData.Token,
		Sign:  authData.Sign,
		Cuit:  authData.Cuit,
	}}
	// Request
	return request[fe.FECompTotXRequest, fe.FECompTotXRequestResponse](structSend)
}
