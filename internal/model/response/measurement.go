package responseModel

import (
	"encoding/json"
	"time"
)

type Measurement struct {
	ID       uint `json:"id,omitempty"`
	IsActive bool `json:"isActive,omitempty"`

	Values json.RawMessage `json:"values,omitempty"`

	PersonId   *uint   `json:"personId,omitempty"`
	Person     *Person `json:"person,omitempty"`
	PersonName string  `json:"personName,omitempty"`

	DressTypeId *uint      `json:"dressTypeId,omitempty"`
	DressType   *DressType `json:"dressType,omitempty"`

	TakenById *uint  `json:"takenById,omitempty"`
	TakenBy   string `json:"takenBy,omitempty"` // first_name + last_name

	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	UpdatedById *uint      `json:"updatedById,omitempty"`
}
