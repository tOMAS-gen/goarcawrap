package certificate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tOMAS-gen/goarcawrap/model"
)

func ViewCSR() (*[]byte,error) {
	// Define la ruta al archivo CSR
	// Idealmente, obtener esto de una constante o función en el paquete certificate
	csrFilePath := filepath.Join(model.DataDir, model.CSRfileName) // Asumiendo que filepath fue importado

	// Leer el contenido del archivo CSR
	csrBytes, err := os.ReadFile(csrFilePath)
	if err != nil {
		// Manejar el caso específico de archivo no encontrado
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Error: El archivo CSR '%s' no fue encontrado. Genera uno primero con '-generate'.\n", csrFilePath)
		} else {
			// Otro error al leer el archivo
			return nil, fmt.Errorf("Error al leer el archivo CSR '%s': %v\n", csrFilePath, err)
		}
	}
	// Devolver el contenido del archivo
	return &csrBytes, nil
}