package wsfev1_model

import "encoding/xml"

// --- Estructuras Comunes de Autenticación y Respuesta ---
const (
	xmlns = "http://ar.gov.afip.dif.FEV1/"
)

// FECompUltimoAutorizado
// Retorna el ultimo comprobante autorizado para el tipo de comprobante / cuit / punto de venta ingresado / Tipo de Emisión
type FECompUltimoAutorizado struct {
	XMLName  xml.Name `xml:"FECompUltimoAutorizado"`
	XMLNS    string   `xml:"xmlns,attr"`
	DataAuth Auth     `xml:"Auth"`
	PtoVta   int      `xml:"PtoVta"`
	CbteTipo int      `xml:"CbteTipo"`
}

func (f *FECompUltimoAutorizado) AddXMLNS() {
	f.XMLNS = xmlns
}

// FECompUltimoAutorizadoResponse
type FECompUltimoAutorizadoResponse struct {
	XMLName                      xml.Name `xml:"FECompUltimoAutorizadoResponse"`
	FECompUltimoAutorizadoResult struct {
		PtoVta   int    `xml:"PtoVta"`
		CbteTipo int    `xml:"CbteTipo"`
		CbteNro  int    `xml:"CbteNro"`
		Errors   *[]Err `xml:"Errors>Err"`
		Events   *[]Evt `xml:"Events>Evt"`
	} `xml:"FECompUltimoAutorizadoResult"`
}

// FEParamGetTiposCbte
// Recupera el listado de Tipos de Comprobantes utilizables en servicio de autorización.
type FEParamGetTiposCbte struct {
	XMLName  xml.Name `xml:"FEParamGetTiposCbte"`
	XMLNS    string   `xml:"xmlns,attr"`
	DataAuth Auth     `xml:"Auth"`
}

func (f *FEParamGetTiposCbte) AddXMLNS() {
	f.XMLNS = xmlns
}

// FEParamGetTiposCbteResponse
type FEParamGetTiposCbteResponse struct {
	XMLName                   xml.Name `xml:"FEParamGetTiposCbteResponse"`
	FEParamGetTiposCbteResult struct {
		ResultGet *[]CbteTipo `xml:"ResultGet>CbteTipo"`
		Errors    *[]Err      `xml:"Errors>Err"`
		Events    *[]Evt      `xml:"Events>Evt"`
	} `xml:"FEParamGetTiposCbteResult"`
}

// FECAEAConsultar
// Consultar CAEA emitidos.
type FECAEAConsultar struct {
	XMLName  xml.Name `xml:"FECAEAConsultar"`
	XMLNS    string   `xml:"xmlns,attr"`
	DataAuth Auth     `xml:"Auth"`
	Periodo  int      `xml:"Periodo"`
	Orden    int16    `xml:"Orden"`
}

func (f *FECAEAConsultar) AddXMLNS() {
	f.XMLNS = xmlns
}

// FECAEAConsultarResponse
type FECAEAConsultarResponse struct {
	XMLName               xml.Name `xml:"FECAEAConsultarResponse"`
	FECAEAConsultarResult struct {
		ResultGet *FECAEAGet `xml:"ResultGet"`
		Errors    *[]Err     `xml:"Errors>Err"`
		Events    *[]Evt     `xml:"Events>Evt"`
	} `xml:"FECAEAConsultarResult"`
}


