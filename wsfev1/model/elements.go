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

// FECAEAGet contiene los datos específicos de un CAEA.
type FECAEAGet struct {
	CAEA          string `xml:"CAEA"`        // Código de Autorización Electrónica de Ambito Ambiental (CAEA)
	Periodo       int    `xml:"Periodo"`     // Periodo del CAEA
	Orden         int16  `xml:"Orden"`       // Orden del CAEA
	FchVigDesde   string `xml:"FchVigDesde"` // Fecha de inicio de vigencia del CAEA
	FchVigHasta   string `xml:"FchVigHasta"` // Fecha de fin de vigencia del CAEA
	FchTopeInf    string `xml:"FchTopeInf"`  // Fecha de tope inferior del CAEA
	FchProceso    string `xml:"FchProceso"`  // Fecha de proceso del CAEA
	Observaciones *[]Obs `xml:"Observaciones>Obs"` // Observaciones del CAEA
}

// Obs representa una observación individual con su código y mensaje.
type Obs struct {
	XMLName xml.Name `xml:"Obs"`
	Code    int      `xml:"Code"` // Código de la observación
	Msg     string   `xml:"Msg"`  // Mensaje descriptivo de la observación
}
