package wsaa

import (
	"fmt"

	"github.com/tOMAS-gen/goarcawrap/model"
	"github.com/tOMAS-gen/goarcawrap/utils"
)

func Save(data *model.WSAA, serviceID string) error {
	return utils.SaveJson(data, fmt.Sprintf("%s-%v", serviceID, model.WSAAfileName))
}
