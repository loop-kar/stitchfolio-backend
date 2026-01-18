package requestModel

type MeasurementHistory struct {
	ID            uint   `json:"id,omitempty"`
	IsActive      bool   `json:"isActive,omitempty"`
	Action        string `json:"action,omitempty"`
	ChangedValues string `json:"changedValues,omitempty"`
	OldValues     string `json:"oldValues,omitempty"` // JSON string
	MeasurementId uint   `json:"measurementId,omitempty"`
	PerformedAt   string `json:"performedAt,omitempty"`
	PerformedById uint   `json:"performedById,omitempty"`
}
