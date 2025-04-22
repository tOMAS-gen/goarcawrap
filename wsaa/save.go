package wsaa

import (
	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/utils"
)

func Save(data *model.WSAA) error {
	return utils.SaveJson(data, model.WSAAfileName)
}
