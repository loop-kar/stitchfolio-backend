package entity

import (
	"testing"

	"github.com/imkarthi24/sf-backend/internal/entities"
)

func TestTemplate(t *testing.T) {
	entities := []interface{}{
		entities.User{},
		entities.Channel{},
		entities.Notification{},
		entities.EmailNotification{},
		entities.WhatsappNotification{},
	}

	ExportTemplate(entities)
}
