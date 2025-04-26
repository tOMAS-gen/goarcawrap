package ws_sr_padrom_a13_api

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap"
	ws_sr_padrom_a13_model "github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13/model"
	ws_sr_padrom_a13_request "github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13/request"
)

type EnvelopeGetPerson struct {
	Body struct {
		Response struct {
			PersonaReturn ws_sr_padrom_a13_model.PersonaReturn `xml:"personaReturn"`
		} `xml:"getPersonaResponse"`
	} `xml:"Body"`
}

func GetPerson(cuit_cuil string) (*ws_sr_padrom_a13_model.PersonaReturn, error) {
	// Obtener Auth
	auth, err := goarcawrap.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	getPerson := ws_sr_padrom_a13_model.GetPersona{
		Token:            auth.Token,
		Sign:             auth.Sign,
		CuitRepresentada: auth.Cuit,
		IdPersona:        cuit_cuil,
	}
	// Crear el xml
	xmlRequest := ws_sr_padrom_a13_request.DefaultXML(getPerson)
	// Enviar la consulta
	response, err := SendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope EnvelopeGetPerson
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.Response.PersonaReturn, nil
}
