package ws_sr_constancia_inscripcion_model

import "encoding/xml"

// Actividad corresponde a tns:actividad
type Actividad struct {
	DescripcionActividad *string  `xml:"descripcionActividad,omitempty"`
	IdActividad          *int64   `xml:"idActividad,omitempty"` // xs:long
	Nomenclador          *int     `xml:"nomenclador,omitempty"` // xs:int
	Orden                *int     `xml:"orden,omitempty"`       // xs:int
	Periodo              *int     `xml:"periodo,omitempty"`     // xs:int
}

// Categoria corresponde a tns:categoria
type Categoria struct {
	DescripcionCategoria *string  `xml:"descripcionCategoria,omitempty"`
	IdCategoria          *int     `xml:"idCategoria,omitempty"` // xs:int
	IdImpuesto           *int     `xml:"idImpuesto,omitempty"`  // xs:int
	Periodo              *int     `xml:"periodo,omitempty"`     // xs:int
}

// Impuesto corresponde a tns:impuesto
type Impuesto struct {
	XMLName             xml.Name `xml:"impuesto"` // Opcional
	DescripcionImpuesto *string  `xml:"descripcionImpuesto,omitempty"`
	EstadoImpuesto      *string  `xml:"estadoImpuesto,omitempty"`
	IdImpuesto          *int     `xml:"idImpuesto,omitempty"` // xs:int
	Motivo              *string  `xml:"motivo,omitempty"`
	Periodo             *int     `xml:"periodo,omitempty"`    // xs:int
}

// --- Definiciones de Tipos para la sección DatosGenerales ---

// Caracterizacion corresponde a tns:caracterizacion
type Caracterizacion struct {
	XMLName                    xml.Name `xml:"caracterizacion"` // Opcional
	DescripcionCaracterizacion *string  `xml:"descripcionCaracterizacion,omitempty"`
	IdCaracterizacion          *int     `xml:"idCaracterizacion,omitempty"` // xs:int
	Periodo                    *int     `xml:"periodo,omitempty"`           // xs:int
}

// Dependencia corresponde a tns:dependencia
type Dependencia struct {
	XMLName                xml.Name `xml:"dependencia"` // Opcional
	CodPostal              *string  `xml:"codPostal,omitempty"`
	DescripcionDependencia *string  `xml:"descripcionDependencia,omitempty"`
	DescripcionProvincia   *string  `xml:"descripcionProvincia,omitempty"`
	Direccion              *string  `xml:"direccion,omitempty"`
	IdDependencia          *int     `xml:"idDependencia,omitempty"` // xs:int
	IdProvincia            *int     `xml:"idProvincia,omitempty"`   // xs:int
	Localidad              *string  `xml:"localidad,omitempty"`
}

// DomicilioA5 corresponde a tns:domicilio (versión de PersonaServiceA5)
type DomicilioA5 struct {
	XMLName              xml.Name `xml:"domicilioFiscal"` // El tag en DatosGenerales es domicilioFiscal
	CodPostal            *string  `xml:"codPostal,omitempty"`
	DatoAdicional        *string  `xml:"datoAdicional,omitempty"`
	DescripcionProvincia *string  `xml:"descripcionProvincia,omitempty"`
	Direccion            *string  `xml:"direccion,omitempty"`
	IdProvincia          *int     `xml:"idProvincia,omitempty"` // xs:int
	Localidad            *string  `xml:"localidad,omitempty"`
	TipoDatoAdicional    *string  `xml:"tipoDatoAdicional,omitempty"`
	TipoDomicilio        *string  `xml:"tipoDomicilio,omitempty"`
}

// DatosGenerales corresponde a tns:datosGenerales
type DatosGenerales struct {
	XMLName             xml.Name           `xml:"datosGenerales"`      // Opcional
	Apellido            *string            `xml:"apellido,omitempty"`
	Caracterizacion     []*Caracterizacion `xml:"caracterizacion,omitempty"` // unbounded nillable
	Dependencia         *Dependencia       `xml:"dependencia,omitempty"`     // minOccurs=0
	DomicilioFiscal     *DomicilioA5       `xml:"domicilioFiscal,omitempty"` // minOccurs=0
	EsSucesion          *string            `xml:"esSucesion,omitempty"`
	EstadoClave         *string            `xml:"estadoClave,omitempty"`
	FechaContratoSocial *string            `xml:"fechaContratoSocial,omitempty"` // xs:dateTime -> *string
	FechaFallecimiento  *string            `xml:"fechaFallecimiento,omitempty"`  // xs:dateTime -> *string
	IdPersona           *int64             `xml:"idPersona,omitempty"`           // xs:long
	MesCierre           *int               `xml:"mesCierre,omitempty"`           // xs:int
	Nombre              *string            `xml:"nombre,omitempty"`
	RazonSocial         *string            `xml:"razonSocial,omitempty"`
	TipoClave           *string            `xml:"tipoClave,omitempty"`
	TipoPersona         *string            `xml:"tipoPersona,omitempty"`
}

// --- Definiciones de Tipos para la sección DatosMonotributo ---

// Relacion corresponde a tns:relacion
type Relacion struct {
	ApellidoPersonaAsociada    *string  `xml:"apellidoPersonaAsociada,omitempty"`
	FfRelacion                 *string  `xml:"ffRelacion,omitempty"`         // xs:dateTime -> *string
	FfVencimiento              *string  `xml:"ffVencimiento,omitempty"`      // xs:dateTime -> *string
	IdPersonaAsociada          *int64   `xml:"idPersonaAsociada,omitempty"`  // xs:long
	NombrePersonaAsociada      *string  `xml:"nombrePersonaAsociada,omitempty"`
	RazonSocialPersonaAsociada *string  `xml:"razonSocialPersonaAsociada,omitempty"`
	TipoComponente             *string  `xml:"tipoComponente,omitempty"`
}

// DatosMonotributo corresponde a tns:datosMonotributo
type DatosMonotributo struct {
	XMLName                 xml.Name    `xml:"datosMonotributo"`              // Opcional
	Actividad               []*Actividad `xml:"actividad,omitempty"`          // unbounded nillable
	ActividadMonotributista *Actividad   `xml:"actividadMonotributista,omitempty"` // minOccurs=0
	CategoriaMonotributo    *Categoria   `xml:"categoriaMonotributo,omitempty"`  // minOccurs=0
	ComponenteDeSociedad    []*Relacion  `xml:"componenteDeSociedad,omitempty"`  // unbounded nillable
	Impuesto                []*Impuesto  `xml:"impuesto,omitempty"`              // unbounded nillable
}

// --- Definiciones de Tipos para la sección DatosRegimenGeneral ---

// Regimen corresponde a tns:regimen
type Regimen struct {
	XMLName            xml.Name `xml:"regimen"` // Opcional
	DescripcionRegimen *string  `xml:"descripcionRegimen,omitempty"`
	IdImpuesto         *int     `xml:"idImpuesto,omitempty"` // xs:int
	IdRegimen          *int     `xml:"idRegimen,omitempty"`  // xs:int
	Periodo            *int     `xml:"periodo,omitempty"`    // xs:int
	TipoRegimen        *string  `xml:"tipoRegimen,omitempty"`
}

// DatosRegimenGeneral corresponde a tns:datosRegimenGeneral
type DatosRegimenGeneral struct {
	XMLName           xml.Name    `xml:"datosRegimenGeneral"` // Opcional
	Actividad         []*Actividad `xml:"actividad,omitempty"` // unbounded nillable
	CategoriaAutonomo *Categoria   `xml:"categoriaAutonomo,omitempty"` // minOccurs=0
	Impuesto          []*Impuesto  `xml:"impuesto,omitempty"` // unbounded nillable
	Regimen           []*Regimen   `xml:"regimen,omitempty"`  // unbounded nillable
}

// --- Definiciones de Tipos para las secciones de Error ---

// ErrorConstancia corresponde a tns:errorConstancia
type ErrorConstancia struct {
	XMLName   xml.Name  `xml:"errorConstancia"`    // Opcional
	Apellido  *string   `xml:"apellido,omitempty"`
	Error     []*string `xml:"error,omitempty"`      // unbounded nillable string
	IdPersona *int64    `xml:"idPersona,omitempty"`  // xs:long
	Nombre    *string   `xml:"nombre,omitempty"`
}

// ErrorMonotributo corresponde a tns:errorMonotributo
type ErrorMonotributo struct {
	XMLName xml.Name  `xml:"errorMonotributo"` // Opcional
	Error   []*string `xml:"error,omitempty"`  // unbounded nillable string
	Mensaje *string   `xml:"mensaje,omitempty"`
}

// ErrorRegimenGeneral corresponde a tns:errorRegimenGeneral
type ErrorRegimenGeneral struct {
	XMLName xml.Name  `xml:"errorRegimenGeneral"` // Opcional
	Error   []*string `xml:"error,omitempty"`   // unbounded nillable string
	Mensaje *string   `xml:"mensaje,omitempty"`
}

// --- Definición de Tipo para Metadata (A5 parece igual a A13) ---

// MetadataA5 corresponde a tns:metadata en PersonaServiceA5
type MetadataA5 struct {
	XMLName   xml.Name `xml:"metadata"`            // Opcional
	FechaHora *string  `xml:"fechaHora,omitempty"` // xs:dateTime -> *string
	Servidor  *string  `xml:"servidor,omitempty"`
}

// --- Estructura Principal: PersonaReturnA5 ---

// PersonaReturnA5 corresponde al complexType tns:personaReturn de PersonaServiceA5
type PersonaReturnA5 struct {
	DatosGenerales      *DatosGenerales      `xml:"datosGenerales,omitempty"`
	DatosMonotributo    *DatosMonotributo    `xml:"datosMonotributo,omitempty"`
	DatosRegimenGeneral *DatosRegimenGeneral `xml:"datosRegimenGeneral,omitempty"`
	ErrorConstancia     *ErrorConstancia     `xml:"errorConstancia,omitempty"`
	ErrorMonotributo    *ErrorMonotributo    `xml:"errorMonotributo,omitempty"`
	ErrorRegimenGeneral *ErrorRegimenGeneral `xml:"errorRegimenGeneral,omitempty"`
	Metadata            *MetadataA5          `xml:"metadata,omitempty"` // Usamos MetadataA5
}

type PersonaListReturnA5 struct {
	XMLName  xml.Name           `xml:"personaListReturn"`         // Tag raíz de la respuesta completa
	Metadata *MetadataA5        `xml:"metadata,omitempty"`        // Metadata general de la respuesta
	Personas []*PersonaReturnA5 `xml:"persona,omitempty"`         // Slice de punteros a PersonaReturnA5
}