package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tOMAS-gen/goarcawrap/model"
)

func ReadJson[T any](nameFile string) (*T, error) {
	// Inicializar
	var data *T

	// Construir la ruta completa al archivo JSON
	filePath := filepath.Join(model.DataDir, nameFile)

	// Verificar si el archivo
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Si el archivo no existe, es normal si aún no se ha guardado.
		return nil, fmt.Errorf("el archivo de cliente %s no existe", filePath)
	} else if err != nil {
		// Otro error al verificar el archivo
		return nil, fmt.Errorf("error al verificar el archivo de cliente %s: %w", filePath, err)
	}

	// Leer el contenido del archivo JSON
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo de cliente %s: %w", filePath, err)
	}

	// Decodificar (deserializar) los datos JSON en una estructura
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON del archivo %s: %w", filePath, err)
	}

	return data, nil // Devuelve el mapa con el cliente y nil como error
}
