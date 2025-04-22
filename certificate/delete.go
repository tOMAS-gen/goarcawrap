package certificate

import (
	"fmt"

	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/util"
)

func DeleteCertificate() error {
	// Eliminar el archivo de clave privada
	err := util.Remove(model.PrivateKeyFileName);
	if err!= nil {
		return fmt.Errorf("error al eliminar la clave privada: %w", err)
	}

	// Eliminar el archivo de certificado
	err = util.Remove(model.CSRfileName);
	if err!= nil {
		return fmt.Errorf("error al eliminar el certificado: %w", err)
	}

	return nil
}
