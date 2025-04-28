package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecaea"
)

func FECAEARegInformativo(feCAEARegInfReq *common.FECAEARequest) (*fecaea.FECAEARegInformativoResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fecaea.FECAEARegInformativo{Auth: *authData,
		FeCAEARegInfReq: feCAEARegInfReq,
	}
	// Request
	return request[fecaea.FECAEARegInformativo, fecaea.FECAEARegInformativoResponse](structSend)
}
