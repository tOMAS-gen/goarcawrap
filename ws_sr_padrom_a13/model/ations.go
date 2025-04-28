package ws_sr_padrom_a13_model

import "encoding/xml"

type GetIdPersonaListByDocumento struct {
	XMLName          xml.Name `xml:"a13:getIdPersonaListByDocumento"`
	Token            string   `xml:"token"`
	Sign             string   `xml:"sign"`
	CuitRepresentada int64    `xml:"cuitRepresentada"`
	Documento        string   `xml:"documento"`
}

type GetPersona struct {
	XMLName          xml.Name `xml:"a13:getPersona"`
	Token            string   `xml:"token"`
	Sign             string   `xml:"sign"`
	CuitRepresentada int64    `xml:"cuitRepresentada"`
	IdPersona        string   `xml:"idPersona"`
}
