package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEParamGetCondicionIvaReceptor(claseCmp *string) (*fe.FEParamGetCondicionIvaReceptorResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FEParamGetCondicionIvaReceptor{Auth: *authData,
		ClaseCmp: claseCmp,
	}
	// Request
	return request[fe.FEParamGetCondicionIvaReceptor, fe.FEParamGetCondicionIvaReceptorResponse](structSend)
}
