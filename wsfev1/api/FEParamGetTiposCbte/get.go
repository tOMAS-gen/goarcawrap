package feparamgettiposcbte

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsaa"
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
	FEParamGetTiposCbteResponse FEParamGetTiposCbteResponse `xml:"FEParamGetTiposCbteResponse"`
}

// FEParamGetTiposCbteResponse
type FEParamGetTiposCbteResponse struct {
	FEParamGetTiposCbteResult wsfev1_model.FEParamGetTiposCbteResult `xml:"FEParamGetTiposCbteResult"`
}

func Get() (*[]wsfev1_model.CbteTipo, error) {
	// Obtener Auth
	auth, err := wsaa.GetAuth(wsfev1_api.ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	getTiposCbte := wsfev1_model.FEParamGetTiposCbte{DataAuth: *auth}
	getTiposCbte.AddXMLNS()
	// Crear el xml
	xmlRequest := wsfev1_request.DefaultXML(getTiposCbte)
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
	return envelope.Body.FEParamGetTiposCbteResponse.FEParamGetTiposCbteResult.ResultGet, nil
}
