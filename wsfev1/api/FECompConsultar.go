package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FECompConsultar(feCompConsReq *common.FECompConsultaReq) (*fe.FECompConsultarResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FECompConsultar{Auth: *authData,
		FeCompConsReq: feCompConsReq,
	}
	// Request
	return request[fe.FECompConsultar, fe.FECompConsultarResponse](structSend)
}
