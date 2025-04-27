package wsfev1_api

import (
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/soap"
)


type Header struct {
	FEHeaderInfo FEHeaderInfo `xml:"FEHeaderInfo"`
}

type FEHeaderInfo struct {
	Ambiente string `xml:"ambiente"`
	Fecha    string `xml:"fecha"`
	ID       string `xml:"id"`
}

const (
	// URL del endpoint del servicio WSFEV1
	urlApi    = model.URL_wsfe1
	ServiceID = "wsfe"
)

func sendSoapRequest(xmlRequestBody string) (string, error) {
	return soap.SendSoapRequest(xmlRequestBody, urlApi)
}
