package requestModel

type SwitchBranch struct {
	EnquiryId uint `json:"enquiryId,omitempty"`
	StudentId uint `json:"studentId,omitempty"`
	ToChannel uint `json:"toChannel,omitempty"`
}
