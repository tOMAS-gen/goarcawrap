package ws_sr_constancia_inscripcion_model

import "encoding/xml"

type GetIdPersonaListByDocumento struct {
	XMLName          xml.Name `xml:"a13:getIdPersonaListByDocumento"`
	Token            string   `xml:"token"`
	Sign             string   `xml:"sign"`
	CuitRepresentada int64    `xml:"cuitRepresentada"`
	Documento        string   `xml:"documento"`
}

type GetPersona struct {
	XMLName          xml.Name `xml:"ns0:getPersona"`
	NS0              string   `xml:"xmlns:ns0,attr"`
	Token            string   `xml:"token"`
	Sign             string   `xml:"sign"`
	CuitRepresentada int64    `xml:"cuitRepresentada"`
	IdPersona        string   `xml:"idPersona"`
}

func NewGetPersona(g *GetPersona) *GetPersona {
	g.NS0 = "http://a5.soap.ws.server.puc.sr/"
	return g
}

type GetPersonaV2 struct {
	XMLName          xml.Name `xml:"ns0:getPersona_v2"`
	NS0              string   `xml:"xmlns:ns0,attr"`
	Token            string   `xml:"token"`
	Sign             string   `xml:"sign"`
	CuitRepresentada int64    `xml:"cuitRepresentada"`
	IdPersona        string   `xml:"idPersona"`
}

func NewGetPersonaV2(g *GetPersonaV2) *GetPersonaV2 {
	g.NS0 = "http://a5.soap.ws.server.puc.sr/"
	return g
}

type GetPersonaListV2 struct {
	XMLName          xml.Name `xml:"ns0:getPersonaList_v2"`
	NS0              string   `xml:"xmlns:ns0,attr"`
	Token            string   `xml:"token"`
	Sign             string   `xml:"sign"`
	CuitRepresentada int64    `xml:"cuitRepresentada"`
	IdPersona        []string `xml:"idPersona"`
}

func NewGetPersonaListV2(g *GetPersonaListV2) *GetPersonaListV2 {
	g.NS0 = "http://a5.soap.ws.server.puc.sr/"
	return g
}
