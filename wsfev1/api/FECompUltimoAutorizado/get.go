package fecompultimoautorizado

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap"
	wsfev1_api "github.com/tOMAS-gen/goarcawrap/wsfev1/api"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
	wsfev1_request "github.com/tOMAS-gen/goarcawrap/wsfev1/request"
)

type Envelope struct {
	Header Header `xml:"Header"`
	Body   Body   `xml:"Body"`
}

type Header struct {
	FEHeaderInfo FEHeaderInfo `xml:"FEHeaderInfo"`
}

type FEHeaderInfo struct {
	Ambiente string `xml:"ambiente"`
	Fecha    string `xml:"fecha"`
	ID       string `xml:"id"`
}

type Body struct {
	FECompUltimoAutorizadoResponse FECompUltimoAutorizadoResponse `xml:"FECompUltimoAutorizadoResponse"`
}

// FECompUltimoAutorizadoResponse
type FECompUltimoAutorizadoResponse struct {
	FECompUltimoAutorizadoResult wsfev1_model.FECompUltimoAutorizadoResult `xml:"FECompUltimoAutorizadoResult"`
}

func Get(ptoVta int, cbteTipo int) (*wsfev1_model.FECompUltimoAutorizadoResult, error) {
	// Obtener Auth
	auth, err := goarcawrap.GetAuth()
	if err != nil {
		return nil, err
	}
	// Crear consulta
	compUltimoAutorizado := wsfev1_model.FECompUltimoAutorizado{DataAuth: *auth, PtoVta: ptoVta, CbteTipo: cbteTipo}
	compUltimoAutorizado.AddXMLNS()
	// Crear el xml
	xmlRequest := wsfev1_request.DefaultXML(compUltimoAutorizado)
	// Enviar la consulta
	response, err := wsfev1_api.SendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope Envelope
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.FECompUltimoAutorizadoResponse.FECompUltimoAutorizadoResult, nil
}
