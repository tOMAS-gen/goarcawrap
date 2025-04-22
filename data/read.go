package data

import (
	"github.com/tOMAS-gen/goarcawrap/model" // Asegúrate que la ruta sea correcta
	"github.com/tOMAS-gen/goarcawrap/utils"
)

// Reload lee el archivo client.json del directorio 'arca'
// y devuelve un mapa conteniendo (potencialmente) ese único cliente.
// Se mantiene el mapa por consistencia, pero podría devolverse *model.Client directamente.
func ReadClient() (*model.Client, error) {
	return utils.ReadJson[model.Client](model.ClientFileName)
}