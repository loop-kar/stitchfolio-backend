package requestModel

type Channel struct {
	ID          uint   `json:"id,omitempty"`
	IsActive    bool   `json:"isActive"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	OwnerUserId uint   `json:"ownerUserId,omitempty"`
}
