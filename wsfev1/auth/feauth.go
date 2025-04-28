package auth

// FEAuthRequest representa la estructura de autenticación
type FEAuthRequest struct {
	Token   string   `xml:"Token,omitempty"`
	Sign    string   `xml:"Sign,omitempty"`
	Cuit    int64    `xml:"Cuit"`
}