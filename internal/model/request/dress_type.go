package requestModel

type DressType struct {
	ID       uint `json:"id,omitempty"`
	IsActive bool `json:"isActive,omitempty"`

	Name         string `json:"name,omitempty"`
	Measurements string `json:"measurements,omitempty"`
}
