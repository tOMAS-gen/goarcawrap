package ws_sr_padrom_a13_request

import (
	"encoding/xml"
)

const (
	// Header is a generic XML header suitable for use with the output of [Marshal].
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	// serviceID es el ID del servicio
	serviceID = "ws_sr_padrom_a13"
)

// EnvelopeSoapEnv representa la estructura principal SOAP 1.2
type EnvelopeSoapEnv struct {
	XMLName xml.Name    `xml:"soapenv:Envelope"`
	SoapEnv string      `xml:"xmlns:soapenv,attr"`
	A13     string      `xml:"xmlns:a13,attr"`
	Body    BodySoapEnv `xml:"soapenv:Body"`
}

func (e *EnvelopeSoapEnv) DefaultEnvelopeSOAP12(body BodySoapEnv) {
	e.SoapEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	e.A13 = "http://a13.soap.ws.server.puc.sr/"
	e.Body = body
}

// BodySoapEnv contiene el cuerpo del mensaje SOAP
type BodySoapEnv struct {
	XMLName xml.Name    `xml:"soapenv:Body"`
	Content interface{} `xml:",any"`
}

func NewXML(content interface{}) string {
	// Generar el XML como []byte
	xmlData, err := xml.MarshalIndent(content, "", "  ")
	if err != nil {
		panic(err)
	}

	// Convertir a string y agregar declaración XML
	return xml.Header + string(xmlData)
}

func DefaultXML(content interface{}) string {
	envelope := EnvelopeSoapEnv{}
	envelope.DefaultEnvelopeSOAP12(BodySoapEnv{
		Content: content,
	})
	return NewXML(envelope)
}
