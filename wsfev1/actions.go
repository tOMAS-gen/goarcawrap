package wsfev1

import (
	wsfev1_api "github.com/tOMAS-gen/goarcawrap/wsfev1/api"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
)

// Web Service orientado al servicio de Facturacion electronica RG2485 V1

// The following operations are supported. For a formal definition, please review the Service Description.

func FECAEAConsultar(periodo int, orden int16) (*wsfev1_model.FECAEAConsultarResponse, error) {
	return wsfev1_api.FECAEAConsultar(periodo, orden)
}

// func FECAEARegInformativo() {

// }

// func FECAEASinMovimientoConsultar() {

// }

// func FECAEASinMovimientoInformar() {

// }

// func FECAEASolicitar() {

// }

// func FECAESolicitar() {

// }

// func FECompConsultar() {

// }

// func FECompTotXRequest() {

// }

// Retorna el ultimo comprobante autorizado para el tipo de comprobante / cuit / punto de venta ingresado / Tipo de Emisión
func FECompUltimoAutorizado(ptoVta int, cbteTipo int) (*wsfev1_model.FECompUltimoAutorizadoResponse, error) {
	return wsfev1_api.FECompUltimoAutorizado(ptoVta, cbteTipo)
}

// func FEDummy() {

// }

// func FEParamGetActividades() {

// }

// func FEParamGetCondicionIvaReceptor() {

// }

// func FEParamGetCotizacion() {

// }

// func FEParamGetPtosVenta() {

// }

// Recupera el listado de Tipos de Comprobantes utilizables en servicio de autorización.
func FEParamGetTiposCbte() (*wsfev1_model.FEParamGetTiposCbteResponse, error) {
	return wsfev1_api.FEParamGetTiposCbte()
}

// func FEParamGetTiposConcepto() {

// }

// func FEParamGetTiposDoc() {

// }

// func FEParamGetTiposIva() {

// }

// func FEParamGetTiposMonedas() {

// }

// func FEParamGetTiposOpcional() {

// }

// func FEParamGetTiposPaises() {

// }

// func FEParamGetTiposTributos() {

// }
