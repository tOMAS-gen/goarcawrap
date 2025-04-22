package data

import (
	"encoding/json" // Para decodificar JSON
	"fmt"           // Para formatear errores
	"os"            // Para leer archivos
	"path/filepath" // Para construir rutas

	"github.com/tOMAS-gen/goarcawrap/model" // Asegúrate que la ruta sea correcta
)

// Reload lee el archivo client.json del directorio 'arca'
// y devuelve un mapa conteniendo (potencialmente) ese único cliente.
// Se mantiene el mapa por consistencia, pero podría devolverse *model.Client directamente.
func ReloadClient() (*model.Client, error) {

	// Inicializar el cliente
	var client *model.Client

	// Construir la ruta completa al archivo client.json en el directorio 'arca'
	filePath := filepath.Join(model.DataDir, model.ClientFileName)

	// Verificar si el archivo client.json existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Si el archivo no existe, es normal si aún no se ha guardado.
		return client, nil // Devuelve mapa vacío, sin error
	} else if err != nil {
		// Otro error al verificar el archivo
		return nil, fmt.Errorf("error al verificar el archivo de cliente %s: %w", filePath, err)
	}

	// Leer el contenido del archivo JSON
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo de cliente %s: %w", filePath, err)
	}

	// Decodificar (deserializar) los datos JSON en una estructura Client
	if err := json.Unmarshal(jsonData, &client); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON del archivo %s: %w", filePath, err)
	}

	return client, nil // Devuelve el mapa con el cliente y nil como error
}