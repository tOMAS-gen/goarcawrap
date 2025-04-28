package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
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
	structSend := fecae.FECAESolicitar{Auth: *authData,
		FeCAEReq: feCAEReq,
	}
	// Request
	return request[fecae.FECAESolicitar, fecae.FECAESolicitarResponse](structSend)
}
