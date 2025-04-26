package ws_sr_padrom_a13_model

import "encoding/xml"

// Domicilio corresponde al complexType tns:domicilio
type Domicilio struct {
	XMLName              xml.Name `xml:"domicilio"` // Opcional, ayuda a la claridad
	Calle                *string  `xml:"calle,omitempty"`
	CodigoPostal         *string  `xml:"codigoPostal,omitempty"`
	DatoAdicional        *string  `xml:"datoAdicional,omitempty"`
	DescripcionProvincia *string  `xml:"descripcionProvincia,omitempty"`
	Direccion            *string  `xml:"direccion,omitempty"`
	EstadoDomicilio      *string  `xml:"estadoDomicilio,omitempty"`
	IdProvincia          *int     `xml:"idProvincia,omitempty"` // xs:int -> int
	Localidad            *string  `xml:"localidad,omitempty"`
	Manzana              *string  `xml:"manzana,omitempty"`
	Numero               *int     `xml:"numero,omitempty"` // xs:int -> int
	OficinaDptoLocal     *string  `xml:"oficinaDptoLocal,omitempty"`
	Piso                 *string  `xml:"piso,omitempty"`
	Sector               *string  `xml:"sector,omitempty"`
	TipoDatoAdicional    *string  `xml:"tipoDatoAdicional,omitempty"`
	TipoDomicilio        *string  `xml:"tipoDomicilio,omitempty"`
	Torre                *string  `xml:"torre,omitempty"`
}

// Metadata corresponde al complexType tns:metadata
type Metadata struct {
	XMLName   xml.Name `xml:"metadata"`            // Opcional
	FechaHora *string  `xml:"fechaHora,omitempty"` // xs:dateTime -> *string
	// FechaHora *time.Time `xml:"fechaHora,omitempty"` // Alternativa con time.Time
	Servidor *string `xml:"servidor,omitempty"`
}

// Persona corresponde al complexType tns:persona
type Persona struct {
	XMLName  xml.Name `xml:"persona"` // Opcional
	Apellido *string  `xml:"apellido,omitempty"`
	// Slice de punteros a int64 para maxOccurs="unbounded" minOccurs="0" nillable="true" type="xs:long"
	ClaveInactivaAsociada         []*int64 `xml:"claveInactivaAsociada,omitempty"`
	DescripcionActividadPrincipal *string  `xml:"descripcionActividadPrincipal,omitempty"`
	// Slice de punteros a Domicilio para maxOccurs="unbounded" minOccurs="0" nillable="true" type="tns:domicilio"
	Domicilio                 []*Domicilio `xml:"domicilio,omitempty"`
	EstadoClave               *string      `xml:"estadoClave,omitempty"`
	FechaContratoSocial       *string      `xml:"fechaContratoSocial,omitempty"` // xs:dateTime -> *string
	FechaFallecimiento        *string      `xml:"fechaFallecimiento,omitempty"`  // xs:dateTime -> *string
	FechaNacimiento           *string      `xml:"fechaNacimiento,omitempty"`     // xs:dateTime -> *string
	FormaJuridica             *string      `xml:"formaJuridica,omitempty"`
	IdActividadPrincipal      *int64       `xml:"idActividadPrincipal,omitempty"` // xs:long -> int64
	IdPersona                 *int64       `xml:"idPersona,omitempty"`            // xs:long -> int64
	MesCierre                 *int         `xml:"mesCierre,omitempty"`            // xs:int -> int
	Nombre                    *string      `xml:"nombre,omitempty"`
	NumeroDocumento           *string      `xml:"numeroDocumento,omitempty"`
	PeriodoActividadPrincipal *int         `xml:"periodoActividadPrincipal,omitempty"` // xs:int -> int
	RazonSocial               *string      `xml:"razonSocial,omitempty"`
	TipoClave                 *string      `xml:"tipoClave,omitempty"`
	TipoDocumento             *string      `xml:"tipoDocumento,omitempty"`
	TipoPersona               *string      `xml:"tipoPersona,omitempty"`
}

// PersonaReturnA13 corresponde al complexType tns:personaReturn
// Este es el struct principal que probablemente necesites deserializar
// desde la respuesta de GetPersona.
type PersonaReturnA13 struct {
	XMLName  xml.Name  `xml:"personaReturn"`      // Opcional, podría venir sin el tag si está dentro de getPersonaResponse
	Metadata *Metadata `xml:"metadata,omitempty"` // minOccurs="0"
	Persona  *Persona  `xml:"persona,omitempty"`  // minOccurs="0"
}
