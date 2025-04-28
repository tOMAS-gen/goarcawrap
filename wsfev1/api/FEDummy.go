package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
)

func FEDummy() (*fe.FEDummyResponse, error) {
	// Datos
	structSend := fe.FEDummy{}
	// Request
	return request[fe.FEDummy, fe.FEDummyResponse ](structSend)
}