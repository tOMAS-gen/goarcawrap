package wsaa

import (
	"fmt"

	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/utils"
)

func read(serviceID string) (*model.WSAA, error) {
	return utils.ReadJson[model.WSAA](fmt.Sprintf("%s-%v", serviceID, model.WSAAfileName))
}
