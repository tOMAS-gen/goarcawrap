package wsfev1_api

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsaa"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
	wsfev1_request "github.com/tOMAS-gen/goarcawrap/wsfev1/request"
)

type FECAEAConsultarResponse struct {
	Header Header `xml:"Header"`
	Body   struct {
		FECAEAConsultarResponse wsfev1_model.FECAEAConsultarResponse `xml:"FECAEAConsultarResponse"`
	} `xml:"Body"`
}

func FECAEAConsultar(periodo int, orden int16) (*wsfev1_model.FECAEAConsultarResponse, error) {
	// Obtener Auth
	auth, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	consultarCAEA := wsfev1_model.FECAEAConsultar{
		DataAuth: *auth,
		Periodo:  periodo,
		Orden:    orden,
	}
	consultarCAEA.AddXMLNS()
	// Crear el xml
	xmlRequest := wsfev1_request.DefaultXML(consultarCAEA)
	// Enviar la consulta
	response, err := sendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope FECAEAConsultarResponse
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.FECAEAConsultarResponse, nil
}
