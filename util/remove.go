package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tOMAS-gen/goarcawrap/model"
)

func Remove(path string) error {
	// Construir la ruta completa al archivo client.json
	filePath := filepath.Join(model.DataDir, path)

	// Intentar eliminar el archivo
	err := os.Remove(filePath)

	// Verificar si hubo un error
	if err != nil {
		// Si el error es porque el archivo no existe, podemos considerarlo éxito
		// (el estado deseado es que el archivo no esté).
		if os.IsNotExist(err) {
			return nil // No es un error real en este contexto
		}
		// Otro tipo de error al intentar eliminar
		return fmt.Errorf("error al eliminar el archivo de cliente %s: %w", filePath, err)
	}

	// Si no hubo error, la eliminación fue exitosa
	return nil // Indica éxito
}
