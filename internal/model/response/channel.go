package responseModel

import "github.com/imkarthi24/sf-backend/internal/entities"

type Channel struct {
	*entities.Channel
	ChannelOwnerFirstName string `json:"channelOwnerFirstName,omitempty"`
	ChannelOwnerLastName  string `json:"channelOwnerLastName,omitempty"`
	PhoneNumber           string `json:"phoneNumber,omitempty"`
	Email                 string `json:"email,omitempty"`
}

type ChannelAutoComplete struct {
	ChannelID uint   `json:"channelId,omitempty"`
	Name      string `json:"name,omitempty"`
}
