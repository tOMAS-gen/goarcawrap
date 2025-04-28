package fe

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsfev1/auth"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
)

type FECompConsultar struct {
	XMLName       xml.Name                  `xml:"http://ar.gov.afip.dif.FEV1/ FECompConsultar"`
	Auth          auth.FEAuthRequest        `xml:"Auth"`
	FeCompConsReq *common.FECompConsultaReq `xml:"FeCompConsReq"`
}

type FECompTotXRequest struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FECompTotXRequest"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
}

type FECompUltimoAutorizado struct {
	XMLName  xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FECompUltimoAutorizado"`
	Auth     auth.FEAuthRequest `xml:"Auth"`
	PtoVta   int                `xml:"PtoVta"`
	CbteTipo int                `xml:"CbteTipo"`
}

type FEDummy struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FEDummy"`
}

type FEParamGetActividades struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetActividades"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
}

type FEParamGetCondicionIvaReceptor struct {
	XMLName  xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetCondicionIvaReceptor"`
	Auth     auth.FEAuthRequest `xml:"Auth"`
	ClaseCmp *string            `xml:"ClaseCmp"`
}

type FEParamGetCotizacion struct {
	XMLName  xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetCotizacion"`
	Auth     auth.FEAuthRequest `xml:"Auth"`
	MonId    *string            `xml:"MonId"`
	FchCotiz *string            `xml:"FchCotiz"`
}

type FEParamGetPtosVenta struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetPtosVenta"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
}

type FEParamGetTiposCbte struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposCbte"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
}
