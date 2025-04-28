// Package goarcawrap simplifica la integración y el consumo de diversos
// Web Services de la Agencia de Recaudación y Control Aduanero (ARCA) de Argentina.
//
// Proporciona una interfaz unificada y fácil de usar para interactuar con servicios
// como WSAA (Autenticación), WSFEV1 (Facturación Electrónica),
// Padrón A5/A13 (Consulta de datos de contribuyentes), entre otros.
//
// El objetivo principal es abstraer la complejidad de las llamadas SOAP,
// el manejo de autenticación (Tokens y Signatures de WSAA), y el parseo
// de las respuestas XML, permitiendo a los desarrolladores integrar estas
// funcionalidades en sus proyectos Go de manera más rápida y sencilla.
package goarcawrap

import (
	"github.com/tOMAS-gen/goarcawrap/functions"
	function_model "github.com/tOMAS-gen/goarcawrap/functions/model"
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion"
	ws_sr_constancia_inscripcion_model "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/model"
	"github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13"
	ws_sr_padrom_a13_model "github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13/model"
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	wsfev1_api "github.com/tOMAS-gen/goarcawrap/wsfev1/api"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/auth"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/common"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fe"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecae"
	"github.com/tOMAS-gen/goarcawrap/wsfev1/fecaea"
)

// Sector de funciones
type a_functions struct {
}

// Método que devuelve el funciones (para encadenamiento)
func Function() *a_functions {
	return &a_functions{}
}

// Obtener datos del cliente
func (a_functions) GetClientData(docCuitCuil string) (*function_model.Client, error) {
	return functions.GetClient(docCuitCuil)
}

//
// ---------------------------------------------------------------------------------------------
//

// Web Service orientado al servicio de Facturacion electronica RG2485 V1
type wsfev1 struct{}

// Método que devuelve el servicio (para encadenamiento)
func WSFEV1() *wsfev1 {
	return &wsfev1{}
}

// FECAEAConsultar
// Consultar CAEA emitidos.
func (wsfev1) FECAEAConsultar(periodo int, orden int16) (*fecaea.FECAEAConsultarResponse, error) {
	return wsfev1_api.FECAEAConsultar(periodo, orden)
}

// FECAEARegInformativo
// Rendición de comprobantes asociados a un CAEA.
func (wsfev1) FECAEARegInformativo(feCAEARegInfReq *common.FECAEARequest) (*fecaea.FECAEARegInformativoResponse, error) {
	return wsfev1_api.FECAEARegInformativo(feCAEARegInfReq)
}

// FECAEASinMovimientoConsultar
// Consulta CAEA informado como sin movimientos.
func (wsfev1) FECAEASinMovimientoConsultar(caea *string, ptoVta int) (*fecaea.FECAEASinMovimientoConsultarResponse, error) {
	return wsfev1_api.FECAEASinMovimientoConsultar(caea, ptoVta)
}

// FECAEASinMovimientoInformar
// Informa CAEA sin movimientos.
func (wsfev1) FECAEASinMovimientoInformar(caea *string, ptoVta int) (*fecaea.FECAEASinMovimientoInformarResponse, error) {
	return wsfev1_api.FECAEASinMovimientoInformar(caea, ptoVta)
}

// FECAEASolicitar
// Solicitud de Código de Autorización Electrónico Anticipado (CAEA)
func (wsfev1) FECAEASolicitar(periodo int, orden int16) (*fecaea.FECAEASolicitarResponse, error) {
	return wsfev1_api.FECAEASolicitar(periodo, orden)
}

// FECAESolicitar
// Solicitud de Código de Autorización Electrónico (CAE)
func (wsfev1) FECAESolicitar(feCAEReq *common.FECAERequest) (*fecae.FECAESolicitarResponse, error) {
	return wsfev1_api.FECAESolicitar(feCAEReq)
}

// FECompConsultar
// Consulta Comprobante emitido y su código.
func (wsfev1) FECompConsultar(feCompConsReq *common.FECompConsultaReq) (*fe.FECompConsultarResponse, error) {
	return wsfev1_api.FECompConsultar(feCompConsReq)
}

// FECompTotXRequest
// Retorna la cantidad maxima de registros que puede tener una invocacion al metodo FECAESolicitar / FECAEARegInformativo
func (wsfev1) FECompTotXRequest() (*fe.FECompTotXRequestResponse, error) {
	return wsfev1_api.FECompTotXRequest()
}

// FECompUltimoAutorizado
// Retorna el ultimo comprobante autorizado para el tipo de comprobante / cuit / punto de venta ingresado / Tipo de Emisión
func (wsfev1) FECompUltimoAutorizado(ptoVta int, cbteTipo int) (*fe.FECompUltimoAutorizadoResponse, error) {
	return wsfev1_api.FECompUltimoAutorizado(ptoVta, cbteTipo)
}

// FEDummy
// Metodo dummy para verificacion de funcionamiento
func (wsfev1) FEDummy() (*fe.FEDummyResponse, error) {
	return wsfev1_api.FEDummy()
}

// FEParamGetActividades
// Recupera el listado de las diferentes actividades habilitadas para el emisor
func (wsfev1) FEParamGetActividades() (*fe.FEParamGetActividadesResponse, error) {
	return wsfev1_api.FEParamGetActividades()
}

// FEParamGetCondicionIvaReceptor
// Recupera la condicion frente al IVA del receptor (para una clase de comprobante determinada o para todos si no se especifica).
func (wsfev1) FEParamGetCondicionIvaReceptor(claseCmp *string) (*fe.FEParamGetCondicionIvaReceptorResponse, error) {
	return wsfev1_api.FEParamGetCondicionIvaReceptor(claseCmp)
}

// FEParamGetCotizacion
// Recupera la cotizacion de la moneda consultada y su fecha
func (wsfev1) FEParamGetCotizacion(monId *string, fchCotiz *string) (*fe.FEParamGetCotizacionResponse, error) {
	return wsfev1_api.FEParamGetCotizacion(monId, fchCotiz)
}

// FEParamGetPtosVenta
// Recupera el listado de puntos de venta registrados y su estado
func (wsfev1) FEParamGetPtosVenta() (*fe.FEParamGetPtosVentaResponse, error) {
	return wsfev1_api.FEParamGetPtosVenta()
}

// FEParamGetTiposCbte
// Recupera el listado de Tipos de Comprobantes utilizables en servicio de autorización.
func (wsfev1) FEParamGetTiposCbte() (*fe.FEParamGetTiposCbteResponse, error) {
	return wsfev1_api.FEParamGetTiposCbte()
}

// FEParamGetTiposConcepto
// Recupera el listado de identificadores para el campo Concepto.
func (wsfev1) FEParamGetTiposConcepto() (*fe.FEParamGetTiposConceptoResponse, error) {
	return wsfev1_api.FEParamGetTiposConcepto()
}

// FEParamGetTiposDoc
// Recupera el listado de Tipos de Documentos utilizables en servicio de autorización.
func (wsfev1) FEParamGetTiposDoc() (*fe.FEParamGetTiposDocResponse, error) {
	return wsfev1_api.FEParamGetTiposDoc()
}

// FEParamGetTiposIva
// Recupera el listado de Tipos de Iva utilizables en servicio de autorización.
func (wsfev1) FEParamGetTiposIva() (*fe.FEParamGetTiposIvaResponse, error) {
	return wsfev1_api.FEParamGetTiposIva()
}

// FEParamGetTiposMonedas
// Recupera el listado de monedas utilizables en servicio de autorización
func (wsfev1) FEParamGetTiposMonedas() (*fe.FEParamGetTiposMonedasResponse, error) {
	return wsfev1_api.FEParamGetTiposMonedas()
}

// FEParamGetTiposOpcional
// Recupera el listado de identificadores para los campos Opcionales
func (wsfev1) FEParamGetTiposOpcional() (*fe.FEParamGetTiposOpcionalResponse, error) {
	return wsfev1_api.FEParamGetTiposOpcional()
}

// FEParamGetTiposPaises
// Recupera el listado de los diferente paises que pueden ser utilizados en el servicio de autorizacion
func (wsfev1) FEParamGetTiposPaises() (*fe.FEParamGetTiposPaisesResponse, error) {
	return wsfev1_api.FEParamGetTiposPaises()
}

// FEParamGetTiposTributos
// Recupera el listado de los diferente tributos que pueden ser utilizados en el servicio de autorizacion
func (wsfev1) FEParamGetTiposTributos() (*fe.FEParamGetTiposTributosResponse, error) {
	return wsfev1_api.FEParamGetTiposTributos()
}

//
// ---------------------------------------------------------------------------------------------
//

// Sector de ws_sr_constancia_inscripcion
type a_ws_sr_constancia_inscripcion struct {
}

// Método que devuelve el servicio (para encadenamiento)
func ConstanciaInscripcion() *a_ws_sr_constancia_inscripcion {
	return &a_ws_sr_constancia_inscripcion{}
}

// Obtener datos
func (a_ws_sr_constancia_inscripcion) GetPerson(cuit_cuil string) (*ws_sr_constancia_inscripcion_model.PersonaReturnA5, error) {
	return ws_sr_constancia_inscripcion.GetPerson(cuit_cuil)
}

// Obtener datos V2
func (a_ws_sr_constancia_inscripcion) GetPersonV2(cuit_cuil string) (*ws_sr_constancia_inscripcion_model.PersonaReturnA5, error) {
	return ws_sr_constancia_inscripcion.GetPersonV2(cuit_cuil)
}

// Obtener lista datos V2
func (a_ws_sr_constancia_inscripcion) GetPersonListV2(cuit_cuil []string) (*ws_sr_constancia_inscripcion_model.PersonaListReturnA5, error) {
	return ws_sr_constancia_inscripcion.GetPersonListV2(cuit_cuil)
}

//
// ---------------------------------------------------------------------------------------------
//

// Sector de ws_sr_padrom_a13
type a_ws_sr_padrom_a13 struct {
}

// Método que devuelve el servicio (para encadenamiento)
func PadromA13() *a_ws_sr_padrom_a13 {
	return & a_ws_sr_padrom_a13{}
}

// Obtener CUIT/CUIL
func (a_ws_sr_padrom_a13) GetCUIT_CUIL(documento string) (*string, error) {
	return ws_sr_padrom_a13.GetCUIT_CUIL(documento)
}

// Obtener datos
func (a_ws_sr_padrom_a13) GetPerson(cuit_cuil string) (*ws_sr_padrom_a13_model.PersonaReturnA13, error) {
	return ws_sr_padrom_a13.GetPerson(cuit_cuil)
}

//
// ---------------------------------------------------------------------------------------------
//

// Sector de WSAA
type a_wsaa struct {
}

// Método que devuelve el servicio (para encadenamiento)
func WSAA() *a_wsaa {
	return &a_wsaa{}
}

// Generar credenciales
func (a_wsaa) GenerateCertificate(client *model.Client) error {
	return wsaa.GenerateCertificate(client)
}

// Obtener credenciales
func (a_wsaa) GetCSR() (*[]byte, error) {
	return wsaa.GetCSR()
}

// Eliminar credenciales
func (a_wsaa) DeleteCertificate() error {
	return wsaa.DeleteCertificate()
}

// Obtener WSAA
func (a_wsaa) GetWSAA(serviceID string) (*model.WSAA, error) {
	return wsaa.GetWSAA(serviceID)
}

// Autenticación WSAA
func (a_wsaa) GetAuth(serviceID string) (*auth.FEAuthRequest, error) {
	return wsaa.GetAuth(serviceID)
}
