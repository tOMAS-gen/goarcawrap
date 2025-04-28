package fecaea

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
)

type FECAEAConsultarResponse struct {
	XMLName               xml.Name                  `xml:"http://ar.gov.afip.dif.FEV1/ FECAEAConsultarResponse"`
	FECAEAConsultarResult *common.FECAEAGetResponse `xml:"FECAEAConsultarResult,omitempty"`
}

type FECAEARegInformativoResponse struct {
	XMLName                    xml.Name               `xml:"http://ar.gov.afip.dif.FEV1/ FECAEARegInformativoResponse"`
	FECAEARegInformativoResult *common.FECAEAResponse `xml:"FECAEARegInformativoResult,omitempty"`
}

type FECAEASinMovimientoConsultarResponse struct {
	XMLName                            xml.Name                         `xml:"http://ar.gov.afip.dif.FEV1/ FECAEASinMovimientoConsultarResponse"`
	FECAEASinMovimientoConsultarResult *common.FECAEASinMovConsResponse `xml:"FECAEASinMovimientoConsultarResult,omitempty"`
}

type FECAEASinMovimientoInformarResponse struct {
	XMLName                           xml.Name                     `xml:"http://ar.gov.afip.dif.FEV1/ FECAEASinMovimientoConsultarResponse"`
	FECAEASinMovimientoInformarResult *common.FECAEASinMovResponse `xml:"FECAEASinMovimientoInformarResult,omitempty"`
}

type FECAEASolicitarResponse struct {
	XMLName               xml.Name                  `xml:"http://ar.gov.afip.dif.FEV1/ FECAEASolicitarResponse"`
	FECAEASolicitarResult *common.FECAEAGetResponse `xml:"FECAEASolicitarResult,omitempty"`
}
