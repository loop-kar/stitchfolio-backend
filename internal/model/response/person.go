package responseModel

type Person struct {
	ID       uint `json:"id,omitempty"`
	IsActive bool `json:"isActive,omitempty"`

	Name string `json:"name,omitempty"`

	CustomerId *uint     `json:"customerId,omitempty"`
	Customer   *Customer `json:"customer,omitempty"`

	Measurements []Measurement `json:"measurements,omitempty"`
}
