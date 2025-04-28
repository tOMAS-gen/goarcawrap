package common

type ArrayOfActividad struct {
	Actividad *[]Actividad `xml:"Actividad"`
}

type ArrayOfCbteAsoc struct {
	CbteAsoc *[]CbteAsoc `xml:"CbteAsoc"`
}

type ArrayOfTributo struct {
	Tributo *[]Tributo `xml:"Tributo"`
}

type ArrayOfAlicIva struct {
	AlicIva *[]AlicIva `xml:"AlicIva"`
}

type ArrayOfOpcional struct {
	Opcional *[]Opcional `xml:"Opcional"`
}

type ArrayOfComprador struct {
	Comprador *[]Comprador `xml:"Comprador"`
}

type ArrayOfFECAEADetRequest struct {
	FECAEADetRequest *[]FECAEADetRequest `xml:"FECAEADetRequest"`
}

type ArrayOfErr struct {
	Err *[]Err `xml:"Err,omitempty"`
}

type ArrayOfEvt struct {
	Evt *[]Evt `xml:"Evt,omitempty"`
}

type ArrayOfObs struct {
	Obs *[]Obs `xml:"Obs,omitempty"`
}

type ArrayOfFECAEADetResponse struct {
	FECAEDetResponse *[]FECAEADetResponse `xml:"FECAEDetResponse"`
}

type ArrayOfFECAEASinMov struct {
	FECAEASinMov *[]FECAEASinMov `xml:"FECAEASinMov"`
}

type ArrayOfFECAEDetRequest struct {
	FECAEDetRequest *[]FECAEDetRequest `xml:"FECAEDetRequest"`
}

type ArrayOfFECAEDetResponse struct {
	FECAEDetResponse *[]FECAEDetResponse `xml:"FECAEDetResponse"`
}

type ArrayOfActividadesTipo struct {
	ActividadesTipo *[]ActividadesTipo `xml:"ActividadesTipo"`
}

type ArrayOfCondicionIvaReceptor struct {
	CondicionIvaReceptor *[]CondicionIvaReceptor `xml:"CondicionIvaReceptor"`
}

type ArrayOfPtoVenta struct {
	PtoVenta *[]PtoVenta `xml:"PtoVenta"`
}

type ArrayOfCbteTipo struct {
	CbteTipo *[]CbteTipo `xml:"CbteTipo"`
}