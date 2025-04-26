package ws_sr_constancia_inscripcion

import (
	"fmt"

	ws_sr_constancia_inscripcion_api "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/api"
	ws_sr_constancia_inscripcion_model "github.com/tOMAS-gen/goarcawrap/ws_sr_constancia_inscripcion/model"
)

func GetPerson(cuit_cuil string) (*ws_sr_constancia_inscripcion_model.PersonaReturnA5, error) {
	if cuit_cuil == "" {
		return nil, fmt.Errorf("El CUIT/CUIL no puede estar vacio")
	}
	return ws_sr_constancia_inscripcion_api.GetPersona(cuit_cuil)
}

func GetPersonV2(cuit_cuil string) (*ws_sr_constancia_inscripcion_model.PersonaReturnA5, error) {
	if cuit_cuil == "" {
		return nil, fmt.Errorf("El CUIT/CUIL no puede estar vacio")
	}
	return ws_sr_constancia_inscripcion_api.GetPersonaV2(cuit_cuil)
}


func GetPersonListV2(cuit_cuil []string) (*ws_sr_constancia_inscripcion_model.PersonaListReturnA5, error) {
	if len(cuit_cuil) == 0 {
		return nil, fmt.Errorf("Los CUIT/CUIL no puede estar vacio")	
	}
	for _, cuit := range cuit_cuil {
		if cuit == "" {
			return nil, fmt.Errorf("Los CUIT/CUIL no puede estar vacio")
		}
	}
	return ws_sr_constancia_inscripcion_api.GetPersonaListV2(cuit_cuil)
}
