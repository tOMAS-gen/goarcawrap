package ws_sr_constancia_inscripcion_api

import (
	"encoding/xml"

	ws_sr_constancia_inscripcion_model "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/model"
	ws_sr_constancia_inscripcion_request "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/request"
	"github.com/tOMAS-gen/goarcawrap/wsaa"
)

type EnvelopeGetPersonaV2 struct {
	Body struct {
		Response struct {
			PersonaReturn ws_sr_constancia_inscripcion_model.PersonaReturnA5 `xml:"personaReturn"`
		} `xml:"getPersona_v2Response"`
	} `xml:"Body"`
}

func GetPersonaV2(cuit_cuil string) (*ws_sr_constancia_inscripcion_model.PersonaReturnA5, error) {
	// Obtener Auth
	auth, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	getPersona := ws_sr_constancia_inscripcion_model.NewGetPersonaV2(&ws_sr_constancia_inscripcion_model.GetPersonaV2{
		Token:            auth.Token,
		Sign:             auth.Sign,
		CuitRepresentada: auth.Cuit,
		IdPersona:        cuit_cuil,
	})
	// Crear el xml
	xmlRequest := ws_sr_constancia_inscripcion_request.DefaultXML(getPersona)
	// Enviar la consulta
	response, err := SendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope EnvelopeGetPersonaV2
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.Response.PersonaReturn, nil
}
