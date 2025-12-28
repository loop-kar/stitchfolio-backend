package util

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	ExpiresAt *time.Time `json:"ExpiresAt"`
	Email     string     `json:"Email"`
	Role      any        `json:"Role"`
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
}

func GenerateJWT(secretKey string, expMinutes int64, values map[string]interface{}) (string, error) {

	mapClaims := jwt.MapClaims{
		"exp": jwt.TimeFunc().Add(time.Duration(expMinutes) * time.Minute).Unix(),
	}

	for key, value := range values {
		mapClaims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "Signing Error", err
	}
	return tokenString, nil
}

func VerifyJWT(token, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errs.NewXError(errs.INVALID, errs.JWT_ERROR, nil)
		}

		return []byte(secretKey), nil
	})
}
