package ws_sr_constancia_inscripcion_api

import (
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/soap"
)

const (
	// URL del endpoint del servicio WSFEV1
	urlApi    = model.URL_ws_sr_constancia_inscripcion
	ServiceID = "ws_sr_constancia_inscripcion"
)

func SendSoapRequest(xmlRequestBody string) (string, error) {
	return soap.SendSoapRequest(xmlRequestBody, urlApi)
}
