package fe

import (
	"encoding/xml"

	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
)

type FECompConsultarResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FECompConsultarResponse"`
	FECompConsultarResult *common.FECompConsultaResponse `xml:"FECompConsultarResult,omitempty"`
}

type FECompTotXRequestResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FECompTotXRequestResponse"`
	FECompTotXRequestResult *common.FERegXReqResponse `xml:"FECompTotXRequestResult,omitempty"`
}

type FECompUltimoAutorizadoResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FECompUltimoAutorizadoResponse"`
	FECompUltimoAutorizadoResult *common.FERecuperaLastCbteResponse `xml:"FECompUltimoAutorizadoResult,omitempty"`
}

type FEDummyResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FEDummyResponse"`
	FEDummyResult *common.DummyResponse `xml:"FEDummyResult,omitempty"`
}

type FEParamGetActividadesResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetActividadesResponse"`
	FEParamGetActividadesResult *common.FEActividadesResponse `xml:"FEParamGetActividadesResult,omitempty"`
}

type FEParamGetCondicionIvaReceptorResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetCondicionIvaReceptorResponse"`
	FEParamGetCondicionIvaReceptorResult *common.CondicionIvaReceptorResponse `xml:"FEParamGetCondicionIvaReceptorResult,omitempty"`
}

type FEParamGetCotizacionResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetCotizacionResponse"`
	FEParamGetCotizacionResult *common.FECotizacionResponse `xml:"FEParamGetCotizacionResult,omitempty"`
}

type FEParamGetPtosVentaResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetPtosVentaResponse"`
	FEParamGetPtosVentaResult *common.FEPtoVentaResponse `xml:"FEParamGetPtosVentaResult,omitempty"`
}

type FEParamGetTiposCbteResponse struct {
	XMLName xml.Name `xml:"http://ar.gov.afip.dif.FEV1/ FEParamGetTiposCbteResponse"`
	FEParamGetTiposCbteResult *common.CbteTipoResponse `xml:"FEParamGetTiposCbteResult,omitempty"`
}
