package functions

import (
	"github.com/tOMAS-gen/goarcawrap/cond_frente_iva"
	function_model "github.com/tOMAS-gen/goarcawrap/functions/model"
	"github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion"
)

func GetCondIVA(cuit_cuil string) (*function_model.CondFrenteIVA, error) {
	// Obtener datos
	data, err := ws_sr_constancia_inscripcion.GetPersonV2(cuit_cuil)
	// Verificar si hay error
	if err != nil {
		return nil, err
	}
	// Obtener condición frente IVA
	return cond_frente_iva.GetCondFrenteIVA(data), nil
}