package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomString(t *testing.T) {
	// Test case 1: Generate password with minimum requirements
	password := GenerateRandomString(12, 2, 2, 4)
	require.Len(t, password, 12)

	// Count special characters
	specialCount := 0
	for _, char := range password {
		if strings.ContainsRune(specialCharSet, char) {
			specialCount++
		}
	}
	require.GreaterOrEqual(t, specialCount, 2)

	// Count numbers
	numCount := 0
	for _, char := range password {
		if strings.ContainsRune(numberSet, char) {
			numCount++
		}
	}
	require.GreaterOrEqual(t, numCount, 2)

	// Count uppercase letters
	upperCount := 0
	for _, char := range password {
		if strings.ContainsRune(upperCharSet, char) {
			upperCount++
		}
	}
	require.GreaterOrEqual(t, upperCount, 4)

	// Test case 2: Generate password with different length
	password = GenerateRandomString(16, 3, 3, 5)
	require.Len(t, password, 16)

	// Test case 3: Generate password with zero minimums
	password = GenerateRandomString(8, 0, 0, 0)
	require.Len(t, password, 8)
}

func TestGeneratePassword(t *testing.T) {
	// Test case 1: Generate default password
	password := GeneratePassword()
	require.Len(t, password, 12)

	// Count special characters
	specialCount := 0
	for _, char := range password {
		if strings.ContainsRune(specialCharSet, char) {
			specialCount++
		}
	}
	require.GreaterOrEqual(t, specialCount, 2)

	// Count numbers
	numCount := 0
	for _, char := range password {
		if strings.ContainsRune(numberSet, char) {
			numCount++
		}
	}
	require.GreaterOrEqual(t, numCount, 2)

	// Count uppercase letters
	upperCount := 0
	for _, char := range password {
		if strings.ContainsRune(upperCharSet, char) {
			upperCount++
		}
	}
	require.GreaterOrEqual(t, upperCount, 4)
}

func TestGenerateRandom(t *testing.T) {
	// Test case 1: Generate random string with length 10
	randomStr := GenerateRandom(10)
	require.Len(t, randomStr, 10)

	// Test case 2: Generate random string with length 20
	randomStr = GenerateRandom(20)
	require.Len(t, randomStr, 20)
}

func TestHashPassword(t *testing.T) {
	// Test case 1: Hash password
	password := "testPassword123"
	secretKey := "testSecretKey"
	hash := HashPassword(password, secretKey)
	require.NotEmpty(t, hash)
	require.NotEqual(t, password, hash)

	// Test case 2: Same password and secret key should produce same hash
	hash2 := HashPassword(password, secretKey)
	require.Equal(t, hash, hash2)

	// Test case 3: Different password should produce different hash
	hash3 := HashPassword("differentPassword", secretKey)
	require.NotEqual(t, hash, hash3)
}

func TestIsPasswordMatching(t *testing.T) {
	// Test case 1: Matching password
	password := "testPassword123"
	secretKey := "testSecretKey"
	hash := HashPassword(password, secretKey)
	require.True(t, IsPasswordMatching(password, hash, secretKey))

	// Test case 2: Non-matching password
	require.False(t, IsPasswordMatching("wrongPassword", hash, secretKey))

	// Test case 3: Different secret key
	require.False(t, IsPasswordMatching(password, hash, "differentSecretKey"))
}

func TestHashPassword2(t *testing.T) {
	//checking negative case -- will fail
	password1 := "Error@404"
	password2 := "Passowrd"

	hashed1 := HashPassword(password1, "secretKey")

	hashed2 := HashPassword(password2, "secretKey")

	require.NotEqual(t, hashed1, hashed2, "two hashes are not identical")

	isMatch := IsPasswordMatching(password1, hashed2, "secretKey")
	require.False(t, isMatch, "Passwords do not match")

}
