package wsfev1

import (
	fecompultimoautorizado "github.com/tOMAS-gen/goarcawrap/wsfev1/api/FECompUltimoAutorizado"
	feparamgettiposcbte "github.com/tOMAS-gen/goarcawrap/wsfev1/api/FEParamGetTiposCbte"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
)

// Web Service orientado al servicio de Facturacion electronica RG2485 V1

// The following operations are supported. For a formal definition, please review the Service Description.

// func FECAEAConsultar() {

// }

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
func FECompUltimoAutorizado(ptoVta int, cbteTipo int) (*wsfev1_model.FECompUltimoAutorizadoResult, error) {
	return fecompultimoautorizado.Get(ptoVta, cbteTipo)
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
func FEParamGetTiposCbte() (*wsfev1_model.FEParamGetTiposCbteResult, error) {
	return feparamgettiposcbte.Get()
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
