package wsaa

import (
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/utils"
)

func Read() (*model.WSAA, error) {
	return utils.ReadJson[model.WSAA](model.WSAAfileName)
}