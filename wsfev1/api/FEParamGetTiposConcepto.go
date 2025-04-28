package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetTiposConcepto() (*fe.FEParamGetTiposConceptoResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetTiposConcepto{Auth: *authData}
	// Request
	return request[fe.FEParamGetTiposConcepto, fe.FEParamGetTiposConceptoResponse](structSend)
}
