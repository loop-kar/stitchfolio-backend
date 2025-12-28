package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/imkarthi24/sf-backend/pkg/entity"
)

func TestPrepareEntityForUpdate(t *testing.T) {

	tstStruct := User{
		IsLoginDisabled:     false,
		IsLoggedIn:          false,
		LastLoginTime:       time.Time{},
		LoginFailureCounter: 12,
		ResetPasswordString: new(string),
	}

	mapVal := PrepareEntityForUpdate(tstStruct)
	fmt.Print(mapVal)

}

func TestGetTypeAttributes(t *testing.T) {

	tstStruct := User{}

	entityName, mapVal := entity.GetEntityAttributes(tstStruct)
	fmt.Print(mapVal)
	fmt.Print(entityName)

}

type User struct {
	*Model
	FirstName           string    `json:"firstName,omitempty"`
	LastName            string    `json:"lastName,omitempty"`
	Extension           string    `gorm:"not null" json:"extension,omitempty"`
	PhoneNumber         string    `gorm:"unique;not null" json:"phoneNumber,omitempty"`
	Email               string    `gorm:"unique;not null" json:"email,omitempty"`
	Password            string    `gorm:"not null" json:"-"`
	IsLoginDisabled     bool      `json:"isLoginDisabled"`
	IsLoggedIn          bool      `json:"isLoggedIn"`
	LastLoginTime       time.Time `json:"lastLoginTime,omitempty"`
	LoginFailureCounter int16     `json:"loginFailureCounter,omitempty"`
	ResetPasswordString *string   `json:"resetPasswordString"`
}

type Model struct {
	ID          uint       `gorm:"primarykey" json:"ID,omitempty"`
	CreatedAt   *time.Time `gorm:"<-:create" json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	IsActive    bool       `gorm:"default:true;type:bool" json:"isActive"`
	CreatedById *uint      `gorm:"<-:create" json:"createdById,omitempty"`
	UpdatedById *uint      `json:"updatedById,omitempty"`
	ChannelId   uint       `gorm:"<-:create" json:"channelId,omitempty"`
}
