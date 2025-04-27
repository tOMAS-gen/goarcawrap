package wsfev1_api

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsaa"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
	wsfev1_request "github.com/tOMAS-gen/goarcawrap/wsfev1/request"
)

type EnvelopeFECompUltimoAutorizadoResponse struct {
	Header Header `xml:"Header"`
	Body   struct {
		FECompUltimoAutorizadoResponse wsfev1_model.FECompUltimoAutorizadoResponse `xml:"FECompUltimoAutorizadoResponse"`
	} `xml:"Body"`
}

func FECompUltimoAutorizado(ptoVta int, cbteTipo int) (*wsfev1_model.FECompUltimoAutorizadoResponse, error) {
	// Obtener Auth
	auth, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	compUltimoAutorizado := wsfev1_model.FECompUltimoAutorizado{DataAuth: *auth, PtoVta: ptoVta, CbteTipo: cbteTipo}
	compUltimoAutorizado.AddXMLNS()
	// Crear el xml
	xmlRequest := wsfev1_request.DefaultXML(compUltimoAutorizado)
	// Enviar la consulta
	response, err := sendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope EnvelopeFECompUltimoAutorizadoResponse
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.FECompUltimoAutorizadoResponse, nil
}
