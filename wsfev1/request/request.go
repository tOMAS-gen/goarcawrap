package wsfev1_request

import (
	"encoding/xml"
)

const (
	// Header is a generic XML header suitable for use with the output of [Marshal].
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// EnvelopeSOAP12 representa la estructura principal SOAP 1.2
type EnvelopeSOAP12 struct {
	XMLName xml.Name   `xml:"soap12:Envelope"`
	XSI     string     `xml:"xmlns:xsi,attr"`
	XSD     string     `xml:"xmlns:xsd,attr"`
	SOAP12  string     `xml:"xmlns:soap12,attr"`
	Body    BodySOAP12 `xml:"soap12:Body"`
}

func (e *EnvelopeSOAP12) DefaultEnvelopeSOAP12(body BodySOAP12) {
	e.XSI = "http://www.w3.org/2001/XMLSchema-instance"
	e.XSD = "http://www.w3.org/2001/XMLSchema"
	e.SOAP12 = "http://www.w3.org/2003/05/soap-envelope"
	e.Body = body
}

// BodySOAP12 contiene el cuerpo del mensaje SOAP
type BodySOAP12 struct {
	XMLName xml.Name    `xml:"soap12:Body"`
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
	envelope := EnvelopeSOAP12{}
	envelope.DefaultEnvelopeSOAP12(BodySOAP12{
		Content: content,
	})
	return NewXML(envelope)
}
