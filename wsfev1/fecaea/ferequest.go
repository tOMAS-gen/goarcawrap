package fecaea

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsfev1/auth"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
)

type FECAEAConsultar struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.dif.FEV1/ FECAEAConsultar"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
	Periodo int                `xml:"Periodo"`
	Orden   int16              `xml:"Orden"`
}

type FECAEARegInformativo struct {
	XMLName         xml.Name              `xml:"http://ar.gov.afip.wsfev1/ FECAEARegInformativo"`
	Auth            auth.FEAuthRequest    `xml:"Auth"`
	FeCAEARegInfReq *common.FECAEARequest `xml:"FeCAEARegInfReq"`
}

type FECAEASinMovimientoConsultar struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.wsfev1/ FECAEASinMovimientoConsultar"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
	CAEA    *string            `xml:"CAEA"`
	PtoVta  int                `xml:"PtoVta"`
}

type FECAEASinMovimientoInformar struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.wsfev1/ FECAEASinMovimientoInformar"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
	CAEA    *string            `xml:"CAEA"`
	PtoVta  int                `xml:"PtoVta"`
}

type FECAEASolicitar struct {
	XMLName xml.Name           `xml:"http://ar.gov.afip.wsfev1/ FECAEASolicitar"`
	Auth    auth.FEAuthRequest `xml:"Auth"`
	Periodo int                `xml:"Periodo"`
	Orden   int16              `xml:"Orden"`
}
