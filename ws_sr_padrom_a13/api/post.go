package ws_sr_padrom_a13_api

import (
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/soap"
)

const (
	// URL del endpoint del servicio WSFEV1
	urlApi    = model.URL_ws_sr_padron_a13
	ServiceID = "ws_sr_padron_a13"
)

func SendSoapRequest(xmlRequestBody string) (string, error) {
	return soap.SendSoapRequest(xmlRequestBody, urlApi)
}
