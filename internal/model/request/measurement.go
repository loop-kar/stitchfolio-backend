package requestModel

import (
	"encoding/json"
)

type Measurement struct {
	ID       uint `json:"id,omitempty"`
	IsActive bool `json:"isActive,omitempty"`

	Values json.RawMessage `json:"values,omitempty"`

	PersonId    *uint `json:"personId,omitempty"`
	DressTypeId *uint `json:"dressTypeId,omitempty"`
	TakenById   *uint `json:"takenById,omitempty"`
}
