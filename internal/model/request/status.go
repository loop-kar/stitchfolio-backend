package requestModel

// Model used for Student and Enquiry status updates
type Status struct {
	Status       string `json:"status,omitempty" binding:"required"`
	StatusReason string `json:"statusReason,omitempty"`
}
