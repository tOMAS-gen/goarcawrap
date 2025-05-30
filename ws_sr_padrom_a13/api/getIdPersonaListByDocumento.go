package ws_sr_padrom_a13_api

import (
	"encoding/xml"

	ws_sr_padrom_a13_model "github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13/model"
	ws_sr_padrom_a13_request "github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13/request"
	"github.com/tOMAS-gen/goarcawrap/wsaa"
)

type EnvelopeGetCUIT_CUIL struct {
	Body struct {
		Response struct {
			IdPersona string `xml:"idPersonaListReturn>idPersona"`
		} `xml:"getIdPersonaListByDocumentoResponse"`
	} `xml:"Body"`
}

func GetCUIT_CUIL(documento string) (*string, error) {
	// Obtener Auth
	auth, err := wsaa.GetAuth(ServiceID)
	if err != nil {
		return nil, err
	}
	// Crear consulta
	getCuit := ws_sr_padrom_a13_model.GetIdPersonaListByDocumento{
		Token:            auth.Token,
		Sign:             auth.Sign,
		CuitRepresentada: auth.Cuit,
		Documento:        documento,
	}
	// Crear el xml
	xmlRequest := ws_sr_padrom_a13_request.DefaultXML(getCuit)
	// Enviar la consulta
	response, err := SendSoapRequest(xmlRequest)
	if err != nil {
		return nil, err
	}
	// Parsear la respuesta
	var envelope EnvelopeGetCUIT_CUIL
	// Parsear el XML en la estructura SoapEnvelope
	err = xml.Unmarshal([]byte(response), &envelope)
	if err != nil {
		return nil, err
	}
	// Devolver la respuesta
	return &envelope.Body.Response.IdPersona, nil
}
