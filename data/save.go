package data

import (
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/utils"
)

func SaveClient(client *model.Client) error {
	return utils.SaveJson(client, model.ClientFileName)
}
