package fecae

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
)

type FECAESolicitarResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FECAESolicitarResponse"`
	FECAESolicitarResult *common.FECAEResponse `xml:"FECAEASolicitarResult,omitempty"`
}