package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/auth"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecae"
)

func FECAESolicitar(feCAEReq *common.FECAERequest) (*fecae.FECAESolicitarResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fecae.FECAESolicitar{Auth: auth.FEAuthRequest{
		Token: authData.Token,
		Sign:  authData.Sign,
		Cuit:  authData.Cuit,
	},
		FeCAEReq: feCAEReq,
	}
	// Request
	return request[fecae.FECAESolicitar, fecae.FECAESolicitarResponse](structSend)
}
