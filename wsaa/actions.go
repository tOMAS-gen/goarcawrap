package wsaa

import (
	"github.com/tOMAS-gen/goarcawrap/certificate"
	"github.com/tOMAS-gen/goarcawrap/data"
	"github.com/tOMAS-gen/goarcawrap/model"
	wsfev1_model "github.com/tOMAS-gen/goarcawrap/wsfev1/model"
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

func GetWSAA(serviceID string) (*model.WSAA, error) {
	return get(serviceID)
}

func GetAuth(serviceID string) (*wsfev1_model.Auth, error) {
	// Obtener WSAA
	wsaa, err := GetWSAA(serviceID)
	if err != nil {
		return nil, err
	}
	// Obtener Usuario
	client, err := data.ReadClient()
	if err != nil {
		return nil, err
	}
	// Generar Auth
	return &wsfev1_model.Auth{
		Token: wsaa.Token,
		Sign:  wsaa.Sign,
		Cuit:  client.CUIT,
	}, nil
}
