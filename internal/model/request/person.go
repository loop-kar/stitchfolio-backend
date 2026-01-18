package requestModel

type Person struct {
	ID       uint `json:"id,omitempty"`
	IsActive bool `json:"isActive,omitempty"`

	Name string `json:"name,omitempty"`

	CustomerId *uint `json:"customerId,omitempty"`
}
