package entities

import "time"

type ExpenseTracker struct {
	*Model `mapstructure:",squash"`

	PurchaseDate *time.Time `json:"purchaseDate,omitempty"`
	BillNumber   string     `json:"billNumber"`
	CompanyName  string     `json:"companyName"`
	Material     string     `json:"material"`
	Price         *float64   `json:"price,omitempty"`
	Location      *string    `json:"location,omitempty"`
	Notes         *string    `json:"notes,omitempty"`
}

func (ExpenseTracker) TableName() string {
	return "stich.ExpenseTrackers"
}

func (ExpenseTracker) TableNameForQuery() string {
	return "\"stich\".\"ExpenseTrackers\" E"
}
