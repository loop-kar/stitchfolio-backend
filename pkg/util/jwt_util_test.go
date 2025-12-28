package util

import (
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestGenerateJWT(t *testing.T) {
	// Test case 1: Generate valid JWT
	secretKey := "testSecretKey"
	expMinutes := int64(60)
	values := map[string]interface{}{
		"email":     "test@example.com",
		"role":      "admin",
		"firstName": "John",
		"lastName":  "Doe",
	}

	token, err := GenerateJWT(secretKey, expMinutes, values)
	require.Nil(t, err)
	require.NotEmpty(t, token)

	// Verify the token
	parsedToken, err := VerifyJWT(token, secretKey)
	require.Nil(t, err)
	require.NotNil(t, parsedToken)

	// Verify claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	require.True(t, ok)
	require.Equal(t, "test@example.com", claims["email"])
	require.Equal(t, "admin", claims["role"])
	require.Equal(t, "John", claims["firstName"])
	require.Equal(t, "Doe", claims["lastName"])

	// Test case 2: Generate JWT with empty values
	token, err = GenerateJWT(secretKey, expMinutes, map[string]interface{}{})
	require.Nil(t, err)
	require.NotEmpty(t, token)

	// Test case 3: Generate JWT with invalid secret key
	token, err = GenerateJWT("", expMinutes, values)
	require.NotNil(t, err)
	require.Equal(t, "Signing Error", token)
}

func TestVerifyJWT(t *testing.T) {
	// Test case 1: Verify valid token
	secretKey := "testSecretKey"
	expMinutes := int64(60)
	values := map[string]interface{}{
		"email": "test@example.com",
	}

	token, err := GenerateJWT(secretKey, expMinutes, values)
	require.Nil(t, err)

	parsedToken, err := VerifyJWT(token, secretKey)
	require.Nil(t, err)
	require.NotNil(t, parsedToken)

	// Test case 2: Verify expired token
	expMinutes = int64(-1) // Expired token
	token, err = GenerateJWT(secretKey, expMinutes, values)
	require.Nil(t, err)

	parsedToken, err = VerifyJWT(token, secretKey)
	require.NotNil(t, err)
	require.Nil(t, parsedToken)

	// Test case 3: Verify with wrong secret key
	token, err = GenerateJWT(secretKey, expMinutes, values)
	require.Nil(t, err)

	parsedToken, err = VerifyJWT(token, "wrongSecretKey")
	require.NotNil(t, err)
	require.Nil(t, parsedToken)

	// Test case 4: Verify invalid token
	parsedToken, err = VerifyJWT("invalid.token.here", secretKey)
	require.NotNil(t, err)
	require.Nil(t, parsedToken)
}
