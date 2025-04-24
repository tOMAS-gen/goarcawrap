package wsfev1_model

import "encoding/xml"

// --- Estructuras Comunes de Autenticación y Respuesta ---
const (
	xmlns = "http://ar.gov.afip.dif.FEV1/"
)

// FEParamGetTiposCbte
// Recupera el listado de Tipos de Comprobantes utilizables en servicio de autorización.
type FEParamGetTiposCbte struct {
	XMLName  xml.Name `xml:"FEParamGetTiposCbte"`
	XMLNS    string   `xml:"xmlns,attr"`
	DataAuth Auth     `xml:"Auth"`
}

func (f *FEParamGetTiposCbte) Create(auth Auth) {
	f.XMLNS = xmlns
	f.DataAuth = auth
}

// FEParamGetTiposCbteResult
type FEParamGetTiposCbteResult struct {
	XMLName   xml.Name    `xml:"FEParamGetTiposCbteResult"`
	ResultGet *[]CbteTipo `xml:"ResultGet>CbteTipo"`
	Errors    *[]Err      `xml:"Errors>Err"`
	Events    *[]Evt      `xml:"Events>Evt"`
}

