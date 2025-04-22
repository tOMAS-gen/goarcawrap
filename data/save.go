package data

import (
	"encoding/json" // Para trabajar con JSON
	"fmt"           // Para formatear errores
	"os"            // Para operaciones del sistema de archivos
	"path/filepath" // Para construir rutas de archivo de forma segura

	"github.com/tOMAS-gen/goarcawrap/model"
)

// NewClient guarda la información del cliente en un archivo JSON específico para él.
func SaveClient(client *model.Client) error {
	// Crear la ruta del directorio específico para el cliente usando su CUIT
	clientDir := filepath.Join(model.DataDir)
  
	// Asegurarse de que el directorio del cliente exista, creándolo si es necesario
	// 0755 son permisos estándar (lectura/ejecución para todos, escritura para el propietario)
	if err := os.MkdirAll(clientDir, 0755); err != nil {
		return fmt.Errorf("error al crear el directorio del cliente %s: %w", clientDir, err)
	}

	// Verificar si el archivo client.json ya existe
	// Si existe, no lo sobrescribimos
	if _, err := os.Stat(filepath.Join(clientDir, model.ClientFileName)); err == nil {
		return fmt.Errorf("el archivo JSON del cliente ya existe en %s", clientDir)
	}

	// Construir la ruta completa al archivo JSON
	filePath := filepath.Join(clientDir, model.ClientFileName)

	// Convertir (serializar) la estructura del cliente a formato JSON
	// Usamos MarshalIndent para que el JSON sea legible (con indentación)
	jsonData, err := json.MarshalIndent(client, "", "  ")
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