package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         int    `json:"age"`
	IsActive    bool   `json:"isActive"`
	PhoneNumber string `json:"phoneNumber"`
}

func TestMapStructFieldsToCamelCase(t *testing.T) {
	// Test case 1: Basic struct mapping
	testStruct := TestStruct{
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		IsActive:    true,
		PhoneNumber: "1234567890",
	}

	result := MapStructFieldsToCamelCase(testStruct)
	require.Equal(t, "FirstName", result["FirstName"])
	require.Equal(t, "LastName", result["LastName"])
	require.Equal(t, "Age", result["Age"])
	require.Equal(t, "IsActive", result["IsActive"])
	require.Equal(t, "PhoneNumber", result["PhoneNumber"])
}

func TestMapCamelStructFieldsToSnakeCase(t *testing.T) {
	// Test case 1: Basic struct mapping
	testStruct := TestStruct{
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		IsActive:    true,
		PhoneNumber: "1234567890",
	}

	result := MapCamelStructFieldsToSnakeCase(testStruct)
	require.Equal(t, "FirstName", result["FirstName"])
	require.Equal(t, "LastName", result["LastName"])
	require.Equal(t, "Age", result["Age"])
	require.Equal(t, "IsActive", result["IsActive"])
	require.Equal(t, "PhoneNumber", result["PhoneNumber"])
}

func TestStructToMap(t *testing.T) {
	// Test case 1: Basic struct to map conversion
	testStruct := TestStruct{
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		IsActive:    true,
		PhoneNumber: "1234567890",
	}

	result := StructToMap(testStruct)
	require.Equal(t, "John", result["firstName"])
	require.Equal(t, "Doe", result["lastName"])
	require.Equal(t, float64(30), result["age"])
	require.Equal(t, true, result["isActive"])
	require.Equal(t, "1234567890", result["phoneNumber"])

	// Test case 2: Empty struct
	emptyStruct := TestStruct{}
	result = StructToMap(emptyStruct)
	require.Equal(t, "", result["firstName"])
	require.Equal(t, "", result["lastName"])
	require.Equal(t, float64(0), result["age"])
	require.Equal(t, false, result["isActive"])
	require.Equal(t, "", result["phoneNumber"])
}
