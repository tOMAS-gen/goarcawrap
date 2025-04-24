package wsfev1_model

import "encoding/xml"

// Auth
type Auth struct {
	XMLName xml.Name `xml:"Auth"`
	Token   string   `xml:"Token,omitempty"`
	Sign    string   `xml:"Sign,omitempty"`
	Cuit    int64    `xml:"Cuit"`
}

// Error
type Err struct {
	XMLName xml.Name `xml:"Err"`
	Code    int      `xml:"Code"`
	Msg     string   `xml:"Msg"`
}

// Evt
type Evt struct {
	XMLName xml.Name `xml:"Evt"`
	Code    int      `xml:"Code"`
	Msg     string   `xml:"Msg"`
}

// CbteTipo
type CbteTipo struct {
	XMLName  xml.Name `xml:"CbteTipo"` // Necesario para el array
	Id       int      `xml:"Id"`       // Código de tipo de comprobante
	Desc     string   `xml:"Desc"`     // Descripción del tipo de comprobante
	FchDesde string   `xml:"FchDesde"` // Fecha desde
	FchHasta *string  `xml:"FchHasta"` // Fecha hasta
}