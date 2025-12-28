package responseModel

import "time"

type Note struct {
	Id               uint       `json:"id,omitempty"`
	FileName         string     `json:"fileName,omitempty"`
	EntityName       string     `json:"entityName,omitempty"`
	EntityIdentifier uint       `json:"entityIdentifier,omitempty"`
	Description      string     `json:"description,omitempty"`
	Type             string     `json:"type,omitempty"`
	Content          string     `json:"content,omitempty"`
	LastModified     *time.Time `json:"lastModified,omitempty"`
	LastModifiedBy   *string    `json:"lastModifiedBy,omitempty"`
}
