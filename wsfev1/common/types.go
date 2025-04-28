package common

import "encoding/xml"

type CbteAsoc struct {
	Tipo    int     `xml:"Tipo"`
	PtoVta  int     `xml:"PtoVta"`
	Nro     int     `xml:"Nro"`
	Cuit    *string `xml:"Cuit"`
	CbteFch *string `xml:"CbteFch"`
}

type Tributo struct {
	Id      int     `xml:"Id"`
	Desc    *string `xml:"Desc"`
	BaseImp float64 `xml:"BaseImp"`
	Alic    float64 `xml:"Alic"`
	Importe float64 `xml:"Importe"`
}

type AlicIva struct {
	Id      int     `xml:"Id"`
	BaseImp float64 `xml:"BaseImp"`
	Importe float64 `xml:"Importe"`
}

type Opcional struct {
	Id    *int    `xml:"Id"`
	Valor *string `xml:"Valor"`
}

type Comprador struct {
	DocTipo int     `xml:"DocTipo"`
	DocNro  int64   `xml:"DocNro"`
	Porcent float64 `xml:"Porcent"`
}

type Periodo struct {
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}

type Actividad struct {
	Id int64 `xml:"Id"`
}

type Err struct {
	XMLName xml.Name `xml:"Err"`
	Code    int      `xml:"Code"`
	Msg     *string  `xml:"Msg,omitempty"`
}

type Evt struct {
	XMLName xml.Name `xml:"Evt"`
	Code    int      `xml:"Code"`
	Msg     *string  `xml:"Msg,omitempty"`
}

type Obs struct {
	XMLName xml.Name `xml:"Obs"`
	Code    int      `xml:"Code"`
	Msg     *string  `xml:"Msg,omitempty"`
}

type DummyResponse struct {
	AppServer  *string `xml:"AppServer"`
	DbServer   *string `xml:"DbServer"`
	AuthServer *string `xml:"AuthServer"`
}

type ActividadesTipo struct {
	Id    int64   `xml:"Id"`
	Orden int16   `xml:"Orden"`
	Desc  *string `xml:"Desc"`
}

type CondicionIvaReceptor struct {
	Id        int     `xml:"Id"`
	Desc      *string `xml:"Desc"`
	Cmp_Clase *string `xml:"Cmp_Clase"`
}

type Cotizacion struct {
	MonId    *int    `xml:"MonId"`
	MonCotiz float64 `xml:"MonCotiz"`
	FchCotiz *string `xml:"FchCotiz"`
}

type PtoVenta struct {
	Nro         int     `xml:"Nro"`
	EmisionTipo *string `xml:"EmisionTipo"`
	Bloqueado   *string `xml:"Bloqueado"`
	FchBaja     *string `xml:"FchBaja"`
}

type CbteTipo struct {
	Id       int     `xml:"Id"`
	Desc     *string `xml:"Desc"`
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}

type ConceptoTipo struct {
	Id       string  `xml:"Id"`
	Desc     *string `xml:"Desc"`
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}

type DocTipo struct {
	Id       int     `xml:"Id"`
	Desc     *string `xml:"Desc"`
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}

type IvaTipo struct {
	Id       *string `xml:"Id"`
	Desc     *string `xml:"Desc"`
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}

type Moneda struct {
	Id       *string `xml:"Id"`
	Desc     *string `xml:"Desc"`
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}

type OpcionalTipo struct {
	Id       *string `xml:"Id"`
	Desc     *string `xml:"Desc"`
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}

type PaisTipo struct {
	Id   int16   `xml:"Id"`
	Desc *string `xml:"Desc"`
}

type TributoTipo struct {
	Id       *string `xml:"Id"`
	Desc     *string `xml:"Desc"`
	FchDesde *string `xml:"FchDesde"`
	FchHasta *string `xml:"FchHasta"`
}
