package common

type FECabRequest struct {
	CantReg  int `xml:"CantReg"`
	PtoVta   int `xml:"PtoVta"`
	CbteTipo int `xml:"CbteTipo"`
}

type FEDetRequest struct {
	Concepto               int               `xml:"Concepto"`
	DocTipo                int               `xml:"DocTipo"`
	DocNro                 int               `xml:"DocNro"`
	CbteDesde              int               `xml:"CbteDesde"`
	CbteHasta              int               `xml:"CbteHasta"`
	CbteFch                *string           `xml:"CbteFch"`
	ImpTotal               float64           `xml:"ImpTotal"`
	ImpTotConc             float64           `xml:"ImpTotConc"`
	ImpNeto                float64           `xml:"ImpNeto"`
	ImpOpEx                float64           `xml:"ImpOpEx"`
	ImpTrib                float64           `xml:"ImpTrib"`
	ImpIVA                 float64           `xml:"ImpIVA"`
	FchServDesde           *string           `xml:"FchServDesde"`
	FchServHasta           *string           `xml:"FchServHasta"`
	FchVtoPago             *string           `xml:"FchVtoPago"`
	MonId                  *string           `xml:"MonId"`
	MonCotiz               *float64          `xml:"MonCotiz"`
	CanMisMonExt           *string           `xml:"CanMisMonExt"`
	CondicionIVAReceptorId *int              `xml:"CondicionIVAReceptorId"`
	CbtesAsoc              *ArrayOfCbteAsoc  `xml:"CbtesAsoc"`
	Tributos               *ArrayOfTributo   `xml:"Tributos"`
	Iva                    *ArrayOfAlicIva   `xml:"Iva"`
	Opcionales             *ArrayOfOpcional  `xml:"Opcionales"`
	Compradores            *ArrayOfComprador `xml:"Compradores"`
	PeriodoAsoc            *Periodo          `xml:"PeriodoAsoc"`
	Actividades            *ArrayOfActividad `xml:"Actividades"`
}

type FECabResponse struct {
	Cuit       int64   `xml:"Cuit"`
	PtoVta     int     `xml:"PtoVta"`
	CbteTipo   int     `xml:"CbteTipo"`
	FchProceso *string `xml:"FchProceso,omitempty"`
	CantReg    int     `xml:"CantReg"`
	Resultado  *string `xml:"Resultado,omitempty"`
	Reproceso  *string `xml:"Reproceso,omitempty"`
}

type FEDetResponse struct {
	Concepto      int         `xml:"Concepto"`
	DocTipo       int         `xml:"DocTipo"`
	DocNro        int64       `xml:"DocNro"`
	CbteDesde     int64       `xml:"CbteDesde"`
	CbteHasta     int64       `xml:"CbteHasta"`
	CbteFch       *string     `xml:"CbteFch,omitempty"`
	Resultado     *string     `xml:"Resultado,omitempty"`
	Observaciones *ArrayOfObs `xml:"Observaciones,omitempty"`
}

type FECompConsultaReq struct {
	CbteTipo int   `xml:"CbteTipo"`
	CbteNro  int64 `xml:"CbteNro"`
	PtoVta   int   `xml:"PtoVta"`
}

type FECompConsultaResponse struct {
	ResultGet *FECompConsResponse `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr         `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt         `xml:"Events,omitempty"`
}

type FECompConsResponse struct {
	FECAEDetRequest
	Resultado       *string     `xml:"Resultado,omitempty"`
	CodAutorizacion *string     `xml:"CodAutorizacion,omitempty"`
	EmisionTipo     *string     `xml:"EmisionTipo,omitempty"`
	FchVto          *string     `xml:"FchVto,omitempty"`
	FchProceso      *string     `xml:"FchProceso,omitempty"`
	Observaciones   *ArrayOfObs `xml:"Observaciones,omitempty"`
	PtoVta          int         `xml:"PtoVta,omitempty"`
	CbteTipo        int         `xml:"CbteTipo,omitempty"`
}

type FERegXReqResponse struct {
	RegXReq int `xml:"RegXReq"`
	Errors  *ArrayOfErr `xml:"Errors,omitempty"`
	Events  *ArrayOfEvt `xml:"Events,omitempty"`
}

type FERecuperaLastCbteResponse struct {
	PtoVta   int `xml:"PtoVta"`
	CbteTipo int `xml:"CbteTipo"`
	CbteNro  int `xml:"CbteNro"`
	Errors   *ArrayOfErr `xml:"Errors,omitempty"`
	Events   *ArrayOfEvt `xml:"Events,omitempty"`
}

type FEActividadesResponse struct {
	ResultGet *ArrayOfActividadesTipo `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr             `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt             `xml:"Events,omitempty"`
}

type CondicionIvaReceptorResponse struct {
	ResultGet *ArrayOfCondicionIvaReceptor `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr                  `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt                  `xml:"Events,omitempty"`
}

type FECotizacionResponse struct {
	ResultGet *Cotizacion `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr        `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt        `xml:"Events,omitempty"`
}

type FEPtoVentaResponse struct {
	ResultGet *ArrayOfPtoVenta `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr      `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt      `xml:"Events,omitempty"`
}

type CbteTipoResponse struct {
	ResultGet *ArrayOfCbteTipo `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr      `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt      `xml:"Events,omitempty"`
}