package common

type FECAEARequest struct {
	FeCabReq *FECAEACabRequest        `xml:"FeCabReq"`
	FeDetReq *ArrayOfFECAEADetRequest `xml:"FeDetReq"`
}

type FECAEACabRequest = FECabRequest

type FECAEADetRequest struct {
	FEDetRequest
	CAEA         *string `xml:"CAEA"`
	CbteFchHsGen *string `xml:"CbteFchHsGen"`
}

type FECAEAGetResponse struct {
	ResultGet *FECAEAGet  `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt `xml:"Events,omitempty"`
}

type FECAEAGet struct {
	CAEA          *string     `xml:"CAEA,omitempty"`
	Periodo       int         `xml:"Periodo"`
	Orden         int16       `xml:"Orden"`
	FchVigDesde   *string     `xml:"FchVigDesde,omitempty"`
	FchVigHasta   *string     `xml:"FchVigHasta,omitempty"`
	FchTopeInf    *string     `xml:"FchTopeInf,omitempty"`
	FchProceso    *string     `xml:"FchProceso,omitempty"`
	Observaciones *ArrayOfObs `xml:"Observaciones,omitempty"`
}

type FECAEAResponse struct {
	FeCabResp *FECAEACabResponse        `xml:"FeCabResp"`
	FeDetResp *ArrayOfFECAEADetResponse `xml:"FeDetResp"`
	Errors    *ArrayOfErr               `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt               `xml:"Events,omitempty"`
}

type FECAEACabResponse = FECabResponse

type FECAEASinMovConsResponse struct {
	ResultGet *ArrayOfFECAEASinMov `xml:"ResultGet,omitempty"`
	Errors    *ArrayOfErr          `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt          `xml:"Events,omitempty"`
}

type FECAEASinMov struct {
	CAEA    *string `xml:"CAEA,omitempty"`
	Periodo *string `xml:"Periodo"`
	Orden   int16   `xml:"Orden"`
}

type FECAEASinMovResponse struct {
	FECAEASinMov
	Resultado *string     `xml:"Resultado"`
	Errors    *ArrayOfErr `xml:"Errors,omitempty"`
	Events    *ArrayOfEvt `xml:"Events,omitempty"`
}

type FECAEADetResponse struct {
	FEDetResponse
	CAEA *string `xml:"CAEA,omitempty"`
}

type FECAEDetResponse struct {
	FEDetResponse
	CAEA      *string `xml:"CAEA,omitempty"`
	CAEFchVto *string `xml:"CAEFchVto,omitempty"`
}
