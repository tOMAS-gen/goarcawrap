package wsfev1_api

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/soap"
	wsfev1_request "github.com/tOMAS-gen/goarcawrap/wsfev1/request"
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

func request[RequestT any, ResponseT any](data RequestT) (*ResponseT, error) {
	// Crear el xml
	xmlRequest := wsfev1_request.DefaultXML(data)
	// Enviar la consulta
	response, err := sendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	type EnvelopeResponse = wsfev1_request.Envelope[ResponseT]
	var envelope EnvelopeResponse
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Retornar
	return &envelope.Body.Content, nil
}
