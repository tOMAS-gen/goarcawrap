package ws_sr_constancia_inscripcion_api

import (
	"encoding/xml"

	ws_sr_constancia_inscripcion_model "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/model"
	ws_sr_constancia_inscripcion_request "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/request"
	"github.com/tOMAS-gen/goarcawrap/wsaa"
)

type EnvelopeGetPersonaListV2 struct {
	Body struct {
		Response struct {
			PersonaListReturn ws_sr_constancia_inscripcion_model.PersonaListReturnA5 `xml:"personaListReturn"`
		} `xml:"getPersonaList_v2Response"`
	} `xml:"Body"`
}

func GetPersonaListV2(cuit_cuil []string) (*ws_sr_constancia_inscripcion_model.PersonaListReturnA5, error) {
	// Obtener Auth
	auth, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	getPersonaList := ws_sr_constancia_inscripcion_model.NewGetPersonaListV2(&ws_sr_constancia_inscripcion_model.GetPersonaListV2{
		Token:            auth.Token,
		Sign:             auth.Sign,
		CuitRepresentada: auth.Cuit,
		IdPersona:        cuit_cuil,
	})
	// Crear el xml
	xmlRequest := ws_sr_constancia_inscripcion_request.DefaultXML(getPersonaList)
	// Enviar la consulta
	response, err := SendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope EnvelopeGetPersonaListV2
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.Response.PersonaListReturn, nil
}
