package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FECompUltimoAutorizado(ptoVta int, cbteTipo int) (*fe.FECompUltimoAutorizadoResponse, error) {
	// Obtener Auth
	authData, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Datos
	structSend := fe.FECompUltimoAutorizado{Auth: *authData,
	PtoVta: ptoVta,
	CbteTipo: cbteTipo,
}
	// Request
	return request[fe.FECompUltimoAutorizado, fe.FECompUltimoAutorizadoResponse](structSend)
}