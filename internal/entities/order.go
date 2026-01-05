package entities

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	*Model `mapstructure:",squash"`

	Status OrderStatus `json:"status"`

	CustomerId *uint     `json:"customerId"`
	Customer   *Customer `gorm:"foreignKey:CustomerId" json:"customer"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderId" json:"orderItems"`
}

func (Order) TableName() string {
	return "stitch.Orders"
}

func (Order) TableNameForQuery() string {
	return "\"stitch\".\"Orders\" E"
}
