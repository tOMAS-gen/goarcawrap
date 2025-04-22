package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tOMAS-gen/goarcawrap/model"
)

func SaveJson[T any](objet *T, nameFile string) error {
	// Crear la ruta del directorio específico para el cliente usando su CUIT
	clientDir := filepath.Join(model.DataDir)

	// Verificar si el archivo ya existe
	// Si existe, no lo sobrescribimos
	if _, err := os.Stat(filepath.Join(clientDir, nameFile)); err == nil {
		return fmt.Errorf("el archivo JSON del cliente ya existe en %s", clientDir)
	}

	// Construir la ruta completa al archivo JSON
	filePath := filepath.Join(clientDir, nameFile)

	// Convertir (serializar) la estructura del cliente a formato JSON
	// Usamos MarshalIndent para que el JSON sea legible (con indentación)
	jsonData, err := json.MarshalIndent(objet, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar los datos del cliente a JSON: %w", err)
	}

	// Escribir los datos JSON al archivo
	// 0644 son permisos estándar (lectura para todos, escritura para el propietario)
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("error al guardar el archivo JSON %s: %w", filePath, err)
	}
	return nil // Indica que la operación fue exitosa
}
