package fe

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
)

type FECompConsultarResponse struct {
	XMLName               xml.Name                       `xml:"http://ar.gov.afip.dif.FEV1/ FECompConsultarResponse"`
	FECompConsultarResult *common.FECompConsultaResponse `xml:"FECompConsultarResult,omitempty"`
}

type FECompTotXRequestResponse struct {
	XMLName                 xml.Name                  `xml:"http://ar.gov.afip.dif.FEV1/ FECompTotXRequestResponse"`
	FECompTotXRequestResult *common.FERegXReqResponse `xml:"FECompTotXRequestResult,omitempty"`
}

type FECompUltimoAutorizadoResponse struct {
	XMLName                      xml.Name                           `xml:"http://ar.gov.afip.dif.FEV1/ FECompUltimoAutorizadoResponse"`
	FECompUltimoAutorizadoResult *common.FERecuperaLastCbteResponse `xml:"FECompUltimoAutorizadoResult,omitempty"`
}

type FEDummyResponse struct {
	XMLName       xml.Name              `xml:"http://ar.gov.afip.dif.FEV1/ FEDummyResponse"`
	FEDummyResult *common.DummyResponse `xml:"FEDummyResult,omitempty"`
}

type FEParamGetActividadesResponse struct {
	XMLName                     xml.Name                      `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetActividadesResponse"`
	FEParamGetActividadesResult *common.FEActividadesResponse `xml:"FEParamGetActividadesResult,omitempty"`
}

type FEParamGetCondicionIvaReceptorResponse struct {
	XMLName                              xml.Name                             `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetCondicionIvaReceptorResponse"`
	FEParamGetCondicionIvaReceptorResult *common.CondicionIvaReceptorResponse `xml:"FEParamGetCondicionIvaReceptorResult,omitempty"`
}

type FEParamGetCotizacionResponse struct {
	XMLName                    xml.Name                     `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetCotizacionResponse"`
	FEParamGetCotizacionResult *common.FECotizacionResponse `xml:"FEParamGetCotizacionResult,omitempty"`
}

type FEParamGetPtosVentaResponse struct {
	XMLName                   xml.Name                   `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetPtosVentaResponse"`
	FEParamGetPtosVentaResult *common.FEPtoVentaResponse `xml:"FEParamGetPtosVentaResult,omitempty"`
}

type FEParamGetTiposCbteResponse struct {
	XMLName                   xml.Name                 `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposCbteResponse"`
	FEParamGetTiposCbteResult *common.CbteTipoResponse `xml:"FEParamGetTiposCbteResult,omitempty"`
}

type FEParamGetTiposConceptoResponse struct {
	XMLName                       xml.Name                     `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposConceptoResponse"`
	FEParamGetTiposConceptoResult *common.ConceptoTipoResponse `xml:"FEParamGetTiposConceptoResult,omitempty"`
}

type FEParamGetTiposDocResponse struct {
	XMLName                  xml.Name                `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposDocResponse"`
	FEParamGetTiposDocResult *common.DocTipoResponse `xml:"FEParamGetTiposDocResult,omitempty"`
}

type FEParamGetTiposIvaResponse struct {
	XMLName                  xml.Name                `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposIvaResponse"`
	FEParamGetTiposIvaResult *common.IvaTipoResponse `xml:"FEParamGetTiposIvaResult,omitempty"`
}

type FEParamGetTiposMonedasResponse struct {
	XMLName                      xml.Name               `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposMonedasResponse"`
	FEParamGetTiposMonedasResult *common.MonedaResponse `xml:"FEParamGetTiposMonedasResult,omitempty"`
}

type FEParamGetTiposOpcionalResponse struct {
	XMLName                       xml.Name                     `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposOpcionalResponse"`
	FEParamGetTiposOpcionalResult *common.OpcionalTipoResponse `xml:"FEParamGetTiposOpcionalResult,omitempty"`
}

type FEParamGetTiposPaisesResponse struct {
	XMLName                     xml.Name               `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposPaisesResponse"`
	FEParamGetTiposPaisesResult *common.FEPaisResponse `xml:"FEParamGetTiposPaisesResult,omitempty"`
}

type FEParamGetTiposTributosResponse struct {
	XMLName                       xml.Name                  `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposTributosResponse"`
	FEParamGetTiposTributosResult *common.FETributoResponse `xml:"FEParamGetTiposTributosResult,omitempty"`
}
