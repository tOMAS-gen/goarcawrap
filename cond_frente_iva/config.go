package cond_frente_iva

import (
	"strings"

	function_model "github.com/tOMAS-gen/goarcawrap/functions/model"
	ws_sr_constancia_inscripcion_model "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/model"
)

// Es muy importante que el orden de los valores sea el mismo que los del ARCA
// por lo tanto, se tiene que verificar que los valores sean los correctos
// y que no se agreguen o eliminen valores sin verificar que sean correctos
var (
	// Relaciones de Impuesto CUIDADO
	relaciones_id_impuesto = map[int]function_model.CondFrenteIVA{
		30: iva_responsable_inscripto,
		32: iva_sujeto_exento,
		34: iva_no_alcanzado,
	}

	// IVA Responsable Inscripto
	iva_responsable_inscripto = function_model.CondFrenteIVA{
		ID:        1,
		Desc:      "IVA Responsable Inscripto",
		Cmp_Clase: []string{"A", "M", "C"},
	} // OK

	// IVA Sujeto Exento
	iva_sujeto_exento = function_model.CondFrenteIVA{
		ID:        4,
		Desc:      "IVA Sujeto Exento",
		Cmp_Clase: []string{"B", "C"},
	} // OK

	// Consumidor Final
	consumidor_final = function_model.CondFrenteIVA{
		ID:        5,
		Desc:      "Consumidor Final",
		Cmp_Clase: []string{"B", "C"},
	}

	// Responsable Monotributo
	responsable_monotributo = function_model.CondFrenteIVA{
		ID:        6,
		Desc:      "Responsable Monotributo",
		Cmp_Clase: []string{"A", "M", "C"},
	} // OK

	// Sujeto No Categorizado
	sujeto_no_categorizado = function_model.CondFrenteIVA{
		ID:        7,
		Desc:      "Sujeto No Categorizado",
		Cmp_Clase: []string{"B", "C"},
	}

	// Proveedor Exterior
	proveedor_exterior = function_model.CondFrenteIVA{
		ID:        8,
		Desc:      "Proveedor del Exterior",
		Cmp_Clase: []string{"B", "C"},
	}

	// Cliente Exterior
	cliente_exterior = function_model.CondFrenteIVA{
		ID:        9,
		Desc:      "Cliente del Exterior",
		Cmp_Clase: []string{"B", "C"},
	}

	// IVA Liberado Ley 19.640
	iva_liberado_ley_19_640 = function_model.CondFrenteIVA{
		ID:        10,
		Desc:      "IVA Liberado - Ley N° 19.640",
		Cmp_Clase: []string{"B", "C"},
	}

	// Monotributista Social
	monotributista_social = function_model.CondFrenteIVA{
		ID:        13,
		Desc:      "Monotributista Social",
		Cmp_Clase: []string{"A", "M", "C"},
	} // OK

	// IVA No Alcanzado
	iva_no_alcanzado = function_model.CondFrenteIVA{
		ID:        15,
		Desc:      "IVA No Alcanzado",
		Cmp_Clase: []string{"B", "C"},
	} // OK

	// Monotributo Trabajador Independiente Promovido
	monotributo_trabajador_independiente_promovido = function_model.CondFrenteIVA{
		ID:        16,
		Desc:      "Monotributo Trabajador Independiente Promovido",
		Cmp_Clase: []string{"A", "M", "C"},
	} // OK
)

func GetCondFrenteIVA(persona *ws_sr_constancia_inscripcion_model.PersonaReturnA5) *function_model.CondFrenteIVA {
	// Verificar si hay datos
	if persona == nil {
		return nil
	}
	// Persona física cuil
	if persona.DatosGenerales == nil {
		return &consumidor_final
	}
	// Verificar si es monotributista
	if persona.DatosMonotributo != nil {
		return verificarCondMonotributo(*persona.DatosMonotributo.CategoriaMonotributo.DescripcionCategoria)
	}
	// Verificar con impuesto
	if persona.DatosRegimenGeneral != nil   {
		return verificarCondImpuesto(persona.DatosRegimenGeneral.Impuesto)
	}
	// Default
	return &consumidor_final
}

func verificarCondMonotributo(categoria string) *function_model.CondFrenteIVA {
	// Verificar SOCIAL
	if strings.Contains(categoria, "SOCIAL") {
		return &monotributista_social
	}
	// Verificar TRABAJADOR INDEPENDIENTE PROMOVIDO
	if strings.Contains(categoria, "PROMOVIDO") {
		return &monotributo_trabajador_independiente_promovido
	}
	// Retornar responsable monotributo
	return &responsable_monotributo
}

func verificarCondImpuesto(impuesto []*ws_sr_constancia_inscripcion_model.Impuesto) *function_model.CondFrenteIVA {
	// Verificar si hay impuesto
	if impuesto == nil  {
		return &consumidor_final
	}
	// Verificar si hay relaciones
	for _, i := range impuesto {
		// Verificar si hay relaciones
		data, exist := relaciones_id_impuesto[*i.IdImpuesto]
		if exist {
			return &data
		}	
	}
	// Retornar nil
	return nil
}
