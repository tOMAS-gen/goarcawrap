package data

import (
	"fmt"

	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/util"
)

// DeleteClient elimina el archivo de datos del cliente (arca/client.json).
func DeleteClient() error {
	err:= util.Remove(model.ClientFileName)
	if err!= nil {
		return fmt.Errorf("error al eliminar el cliente: %w", err)
	}
	return nil
}
