package model

import "github.com/imkarthi24/sf-backend/internal/entities"

type EnquiryExtension struct {
	*entities.Enquiry `mapstructure:",squash"`
	StudentStatus     string `json:"studentStatus,omitempty"`
}
