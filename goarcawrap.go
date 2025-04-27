package goarcawrap

import (
	"github.com/tOMAS-gen/goarcawrap/functions"
	function_model "github.com/tOMAS-gen/goarcawrap/functions/model"
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/wsaa"
	"github.com/tOMAS-gen/goarcawrap/wsfev1"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
)

// Sector de funciones
type a_functions struct {
}

// Funcione
func Function() a_functions {
	// Aquí podrías obtener datos de algún lado
	info := a_functions{}
	return info
}

// Obtener datos del cliente
func (a_functions) GetClientData(docCuitCuil string) (*function_model.Client, error) {
	return functions.GetClient(docCuitCuil)
}

//
// ---------------------------------------------------------------------------------------------
//

// Sector de funciones
type a_wsfev1 struct {
}

// WSFFEV1
func WSFFEV1() a_wsfev1 {
	// Aquí podrías obtener datos de algún lado
	info := a_wsfev1{}
	return info
}

// Consultar CAEA emitidos.
func (a_wsfev1) SearchCAEA(periodo int, orden int16) (*wsfev1_model.FECAEAConsultarResponse, error) {
	return wsfev1.FECAEAConsultar(periodo, orden)
}

// Obtener tipos de comprobantes
func (a_wsfev1) GetTiposCbte() (*wsfev1_model.FEParamGetTiposCbteResponse, error) {
	return wsfev1.FEParamGetTiposCbte()
}

// Ultimo comprobante autorizado
func (a_wsfev1) GetUltimoCompAutorizado(ptoVta int, cbteTipo int) (*wsfev1_model.FECompUltimoAutorizadoResponse, error) {
	return wsfev1.FECompUltimoAutorizado(ptoVta, cbteTipo)
}


//
// ---------------------------------------------------------------------------------------------
//

// Sector de WSAA
type a_wsaa struct {
}

// wsaa
func WSAA() a_wsaa {
	// Aquí podrías obtener datos de algún lado
	info := a_wsaa{}
	return info
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
func (a_wsaa) GetAuth(serviceID string) (*wsfev1_model.Auth, error) {
	return wsaa.GetAuth(serviceID)
}
