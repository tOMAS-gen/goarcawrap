package common

type FECAERequest struct {
	FeCabReq *FECAECabRequest        `xml:"FeCabReq"`
	FeDetReq *ArrayOfFECAEDetRequest `xml:"FeDetReq"`
}

type FECAECabRequest = FECabRequest

type FECAEDetRequest = FEDetRequest

type FECAEResponse struct {
	FeCabResp *FECAECabResponse        `xml:"FeCabResp"`
	FeDetResp *ArrayOfFECAEDetResponse `xml:"FeDetResp"`
	Events    *ArrayOfEvt      `xml:"Events"`
	Errors    *ArrayOfErr       `xml:"Errors"`
}

type FECAECabResponse = FECabResponse


