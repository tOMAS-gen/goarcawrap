// goarcawrap.go
package goarcawrap

import (
	"github.com/tOMAS-gen/goarcawrap/certificate"
	"github.com/tOMAS-gen/goarcawrap/data"
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/wsaa"
)

func GenerateCertificate(client *model.Client) error {
	privateKey, err := certificate.GenerateKeyRSA()
	if err != nil {
		return err
	}
	err = certificate.GenerateCSR(client, privateKey)
	if err != nil {
		return err
	}
	err = data.SaveClient(client)
	if err != nil {
		return err
	}
	return nil
}

func GetCSR() (*[]byte, error) {
	return certificate.ViewCSR()
}

func DeleteCertificate() error {
	if err := data.DeleteClient(); err != nil {
		return err
	}
	if err := certificate.DeleteCertificate(); err != nil {
		return err
	}
	return nil
}

func GetWSAA() (*model.WSAA, error) {
	return wsaa.Get()
}
