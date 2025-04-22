package model

const (
	DataDir            = "ARCA"               // Directorio base para guardar datos de clientes
	PrivateKeyFileName = "MiClavePrivada.key" // Nombre para el archivo de clave privada
	CSRfileName        = "MiPedido.csr"       // Nombre para el archivo CSR
	CRTfileName        = "MiCertificado.crt"  // Nombre para el archivo CSR
	KeyBits            = 2048                 // Tamaño de la clave RSA (mínimo 2048 para AFIP/ARCA)
	ClientFileName     = "client.json"        // Nombre del archivo JSON
)

type Client struct {
	CUIT             string `json:"cuit"`
	CommonName       string `json:"commonName"`
	OrganizationName string `json:"organizationName"`
}
