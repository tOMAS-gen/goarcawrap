package fecae

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsfev1/auth"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
)

type FECAESolicitar struct {
	XMLName  xml.Name             `xml:"http://ar.gov.afip.dif.FEV1/ FECAEAConsultar"`
	Auth     auth.FEAuthRequest   `xml:"Auth"`
	FeCAEReq *common.FECAERequest `xml:"FeCAEReq"`
}
