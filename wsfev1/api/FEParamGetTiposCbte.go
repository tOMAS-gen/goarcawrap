package wsfev1_api

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsaa"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
	wsfev1_request "github.com/tOMAS-gen/goarcawrap/wsfev1/request"
)

type EnvelopeFEParamGetTiposCbteResponse struct {
	Header Header `xml:"Header"`
	Body   struct {
		FEParamGetTiposCbteResponse wsfev1_model.FEParamGetTiposCbteResponse `xml:"FEParamGetTiposCbteResponse"`
	} `xml:"Body"`
}

func FEParamGetTiposCbte() (*wsfev1_model.FEParamGetTiposCbteResponse, error) {
	// Obtener Auth
	auth, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	getTiposCbte := wsfev1_model.FEParamGetTiposCbte{DataAuth: *auth}
	getTiposCbte.AddXMLNS()
	// Crear el xml
	xmlRequest := wsfev1_request.DefaultXML(getTiposCbte)
	// Enviar la consulta
	response, err := sendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope EnvelopeFEParamGetTiposCbteResponse
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.FEParamGetTiposCbteResponse, nil
}
