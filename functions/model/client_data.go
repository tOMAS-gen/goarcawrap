package function_model

import (
	"strconv"
	"strings"
)

type Client struct {
	Persona       Persona
	Nombre        *string
	Apellido      *string
	RazonSocial   *string
	Domicilio     *Domicilio
	CondFrenteIVA *CondFrenteIVA
}

type Persona struct {
	ID        int64
	TipoClave string
	Tipo      string
}

func (p Persona) Data() string {
	return p.TipoClave + " " + strconv.FormatInt(p.ID, 10) + " " + p.Tipo
}

type Domicilio struct {
	Direccion     *string
	CodPostal     *string
	Localidad     *string
	Provincia     *string
	DatoAdicional *string
}

func (d Domicilio) Data() string {
	return *d.Direccion + " " + *d.CodPostal + " " + *d.Localidad + " " + *d.Provincia
}

type CondFrenteIVA struct {
	ID        int64
	Desc      string
	Cmp_Clase []string
}

func (c CondFrenteIVA) Data() string {
	return c.Desc + " " +  strings.Join(c.Cmp_Clase, "/")
}

func (c Client) Name() string {
	if c.RazonSocial == nil {
		return string(*c.Apellido + " " + *c.Nombre)
	}
	return *c.RazonSocial
}
