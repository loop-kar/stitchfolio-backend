package requestModel

type OrderHistory struct {
	ID                   uint    `json:"id,omitempty"`
	IsActive             bool    `json:"isActive,omitempty"`
	Action               string  `json:"action,omitempty"`
	ChangedFields        string  `json:"changedFields,omitempty"`
	Status               *string `json:"status,omitempty"`
	ExpectedDeliveryDate *string `json:"expectedDeliveryDate,omitempty"`
	DeliveredDate        *string `json:"deliveredDate,omitempty"`
	OrderItemId          *uint   `json:"orderItemId,omitempty"`
	OrderItemData        string  `json:"orderItemData,omitempty"` // JSON string
	OrderId              uint    `json:"orderId,omitempty"`
	PerformedAt          string  `json:"performedAt,omitempty"`
	PerformedById        uint    `json:"performedById,omitempty"`
}
