package ws_sr_padrom_a13

import (
	"fmt"

	ws_sr_padrom_a13_api "github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13/api"
	ws_sr_padrom_a13_model "github.com/tOMAS-gen/goarcawrap/ws_sr_padrom_a13/model"
)

func GetCUIT(documento string) (*string, error) {
	if documento == "" {
		return nil, fmt.Errorf("El documento no puede estar vacio")
	}
	return ws_sr_padrom_a13_api.GetCUIT_CUIL(documento)
}

func GetPerson(cuit_cuil string) (*ws_sr_padrom_a13_model.PersonaReturnA13, error) {
	if cuit_cuil == "" {
		return nil, fmt.Errorf("El CUIT/CUIL no puede estar vacio")
	}
	return ws_sr_padrom_a13_api.GetPerson(cuit_cuil)
}
