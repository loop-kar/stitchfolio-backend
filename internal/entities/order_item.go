package entities

type OrderItem struct {
	*Model `mapstructure:",squash"`

	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Total       float64 `json:"total"`

	OrderId uint   `json:"orderId"`
	Order   *Order `gorm:"foreignKey:OrderId" json:"order"`
}

func (OrderItem) TableName() string {
	return "stitch.OrderItems"
}

func (OrderItem) TableNameForQuery() string {
	return "\"stitch\".\"OrderItems\" E"
}
