package responseModel

import "time"

type ExpenseTracker struct {
	ID       uint `json:"id,omitempty"`
	IsActive bool `json:"isActive,omitempty"`

	PurchaseDate *time.Time `json:"purchaseDate,omitempty"`
	BillNumber   string     `json:"billNumber,omitempty"`
	CompanyName  string     `json:"companyName,omitempty"`
	Material     string     `json:"material,omitempty"`
	Price         *float64   `json:"price,omitempty"`
	Location      *string    `json:"location,omitempty"`
	Notes         *string    `json:"notes,omitempty"`

	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	CreatedById *uint      `json:"createdById,omitempty"`
	UpdatedById *uint      `json:"updatedById,omitempty"`
}
