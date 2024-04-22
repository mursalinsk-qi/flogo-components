package customEvent

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	DeviceId string      `md:"device_id"`
	Data     interface{} `md:"data"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"device_id": i.DeviceId,
		"data":      i.Data,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.DeviceId, err = coerce.ToString(values["device_id"])
	if err != nil {
		return err
	}
	i.Data = values["data"]

	return nil
}
