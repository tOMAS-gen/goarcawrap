package functions

import (
	"strings"

	"github.com/tOMAS-gen/goarcawrap/cond_frente_iva"
	function_model "github.com/tOMAS-gen/goarcawrap/functions/model"
	"github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion"
	ws_sr_constancia_inscripcion_model "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/model"
	"github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13"
)

func GetClient(docCuitCuil string) (*function_model.Client, error) {
	// Variables
	var err error
	var data *ws_sr_constancia_inscripcion_model.PersonaReturnA5
	// Obtener datos
	data, err = ws_sr_constancia_inscripcion.GetPersonV2(docCuitCuil)
	if err == nil {
		// Si no hay error, retornar
		return parceClient(*data), nil
	}
	// Verificar si es un error de tipo "No existe persona con ese Id"
	if strings.Contains(err.Error(), "No existe persona con ese Id") {
		// Variables
		var cuit_cuil *string
		// Obtener cuit o cuil
		cuit_cuil, err = ws_sr_padrom_a13.GetCUIT_CUIL(docCuitCuil)
		if err != nil {
			return nil, err
		}
		// Obtener datos
		data, err = ws_sr_constancia_inscripcion.GetPersonV2(*cuit_cuil)
		if err == nil {
			// Si no hay error, retornar
			return parceClient(*data), nil
		}
	}
	// Si es otro error, retornar
	return nil, err
}

func parceClient(data ws_sr_constancia_inscripcion_model.PersonaReturnA5) *function_model.Client {
	// Variables
	client := function_model.Client{}
	// Name
	addName(&client, &data)
	// Tipo
	addTipo(&client, &data)
	// Domicilio
	addDomicilio(&client, &data)
	// Condición frente al IVA
	addCondFrenteIVA(&client, &data)
	// Retornar
	return &client
}

func addName(client *function_model.Client, data *ws_sr_constancia_inscripcion_model.PersonaReturnA5) {
	// Cuenta no activa
	if data.DatosGenerales == nil {
		// Nombre y Apellido
		client.Apellido = data.ErrorConstancia.Apellido
		client.Nombre = data.ErrorConstancia.Nombre
		return
	}
	// Cuenta activa
	if data.DatosGenerales.RazonSocial == nil {
		// Nombre y Apellido
		client.Apellido = data.DatosGenerales.Apellido
		client.Nombre = data.DatosGenerales.Nombre
		return
	}
	// Razón social
	client.RazonSocial = data.DatosGenerales.RazonSocial
}

func addTipo(client *function_model.Client, data *ws_sr_constancia_inscripcion_model.PersonaReturnA5) {
	// No hay datos
	if data.DatosGenerales == nil {
		// Datos persona
		client.Persona = function_model.Persona{
			ID:        *data.ErrorConstancia.IdPersona,
			TipoClave: "CUIL",
			Tipo:      "FISICA",
		}
		return
	}
	// Datos persona
	client.Persona = function_model.Persona{
		ID:        *data.DatosGenerales.IdPersona,
		TipoClave: *data.DatosGenerales.TipoClave,
		Tipo:      *data.DatosGenerales.TipoPersona,
	}
}

func addDomicilio(client *function_model.Client, data *ws_sr_constancia_inscripcion_model.PersonaReturnA5) {
	// No hay datos
	if data.DatosGenerales == nil {
		return
	}
	// Domicilio
	if data.DatosGenerales.DomicilioFiscal != nil {
		client.Domicilio = &function_model.Domicilio{
			Direccion:     data.DatosGenerales.DomicilioFiscal.Direccion,
			CodPostal:     data.DatosGenerales.DomicilioFiscal.CodPostal,
			Localidad:     data.DatosGenerales.DomicilioFiscal.Localidad,
			Provincia:     data.DatosGenerales.DomicilioFiscal.DescripcionProvincia,
			DatoAdicional: data.DatosGenerales.DomicilioFiscal.DatoAdicional,
		}
	}
}

func addCondFrenteIVA(client *function_model.Client, data *ws_sr_constancia_inscripcion_model.PersonaReturnA5) {
	// Condición frente al IVA
	client.CondFrenteIVA = cond_frente_iva.GetCondFrenteIVA(data)
}
