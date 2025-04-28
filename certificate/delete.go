package certificate

import (
	"fmt"

	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/utils"
)

func DeleteCertificate() error {
	// Eliminar el archivo de clave privada
	err := utils.Remove(model.PrivateKeyFileName)
	if err != nil {
		return fmt.Errorf("error al eliminar la clave privada: %w", err)
	}

	// Eliminar el archivo de certificado
	err = utils.Remove(model.CSRfileName)
	if err != nil {
		return fmt.Errorf("error al eliminar el certificado: %w", err)
	}

	return nil
}
