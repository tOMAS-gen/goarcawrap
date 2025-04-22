package certificate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath" // Importar para manejo de rutas

	"github.com/tOMAS-gen/goarcawrap/model"
)

// GenerateKeyRSA genera una clave privada RSA y la guarda en formato PEM dentro del directorio 'arca'.
func GenerateKeyRSA() (*rsa.PrivateKey, error) {

	// --- Verificar si ya existe una clave privada ---
	privKeyPath := filepath.Join(model.DataDir, model.PrivateKeyFileName)
	if _, err := os.Stat(privKeyPath); err == nil {
		return nil, fmt.Errorf("private key already exists at %s", privKeyPath)
	}

	// --- Generar Clave Privada ---
	privateKey, err := rsa.GenerateKey(rand.Reader, model.KeyBits)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	// Asegurarse de que el directorio de salida exista
	if err = os.MkdirAll(model.DataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory %s: %w", model.DataDir, err)
	}

	// --- Guardar Clave Privada en archivo PEM ---
	privKeyFile, err := os.Create(privKeyPath) // Usar la ruta completa
	if err != nil {
		return nil, fmt.Errorf("failed to create private key file: %w", err)
	}
	defer privKeyFile.Close() // Asegura que el archivo se cierre

	// Codificar la clave privada en formato PEM
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey), // Usar PKCS#1 standard
	}
	if err := pem.Encode(privKeyFile, privateKeyPEM); err != nil {
		return nil, fmt.Errorf("failed to write private key PEM data: %w", err)
	}

	// Devolvemos la clave privada generada para poder usarla después (ej: para el CSR)
	return privateKey, nil
}

// GenerateCSR genera un CSR y lo guarda en formato PEM dentro del directorio 'arca'.
func GenerateCSR(client *model.Client, privateKey *rsa.PrivateKey) error {

	// --- Verificar si ya existe un CSR ---
	csrPath := filepath.Join(model.DataDir, model.CSRfileName)
	if _, err := os.Stat(csrPath); err == nil {
		return fmt.Errorf("CSR already exists at %s", csrPath)
	}

	// Datos del subject
	subject := pkix.Name{
		Country:      []string{"AR"},                    // C=AR (o client.CountryCode si existe)
		Organization: []string{client.OrganizationName}, // O=subj_o
		CommonName:   client.CommonName,                 // CN=subj_cn
		SerialNumber: "CUIT " + client.CUIT,                 // serialNumber=CUIT subj_cuit
	}

	// Plantilla para el CSR
	csrTemplate := x509.CertificateRequest{
		Subject:            subject,
		SignatureAlgorithm: x509.SHA256WithRSA, // Algoritmo de firma recomendado
	}

	// Crear el CSR firmándolo con la clave privada
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create CSR: %w", err)
	}

	// Asegurarse de que el directorio de salida exista (por si GenerateCSR se llama independientemente)
	if err = os.MkdirAll(model.DataDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", model.DataDir, err)
	}

	// --- Guardar CSR en archivo PEM ---
	csrFile, err := os.Create(csrPath) // Usar la ruta completa
	if err != nil {
		return fmt.Errorf("failed to create CSR file: %w", err)
	}
	defer csrFile.Close()
  
	// Codificar el CSR en formato PEM
	csrPEM := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	}
	if err := pem.Encode(csrFile, csrPEM); err != nil {
		return fmt.Errorf("failed to write CSR PEM data: %w", err)
	}

	return nil // Indica que la operación fue exitosa
}
