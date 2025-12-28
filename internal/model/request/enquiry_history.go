package requestModel

type EnquiryHistory struct {
	ID              uint    `json:"id,omitempty"`
	IsActive        bool    `json:"isActive,omitempty"`
	EmployeeComment string  `json:"employeeComment,omitempty"`
	CustomerComment string  `json:"customerComment,omitempty"`
	VisitingDate    *string `json:"visitingDate,omitempty"`
	CallBackDate    *string `json:"callBackDate,omitempty"`
	EnquiryDate     string  `json:"enquiryDate,omitempty"`
	ResponseStatus  string  `json:"responseStatus,omitempty"`
	EnquiryId       uint    `json:"enquiryId,omitempty"`
	EmployeeId      uint    `json:"employeeId,omitempty"`
}
