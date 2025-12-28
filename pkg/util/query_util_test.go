package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildQuery(t *testing.T) {
	// Test case 1: Basic query with single filter
	filters := "id eq 1"
	fieldMap := map[string]string{
		"Id": "id",
	}
	expected := "id = 1"
	result := BuildQuery(filters, fieldMap)
	require.Equal(t, expected, result)

	// Test case 2: Multiple filters
	filters = "id eq 1,amount gt 1000"
	fieldMap = map[string]string{
		"Id":     "id",
		"Amount": "amount",
	}
	expected = "id = 1 AND amount > 1000"
	result = BuildQuery(filters, fieldMap)
	require.Equal(t, expected, result)

	// Test case 3: Filter with non-existent field
	filters = "id eq 1,invalidField eq value"
	fieldMap = map[string]string{
		"Id": "id",
	}
	expected = "id = 1"
	result = BuildQuery(filters, fieldMap)
	require.Equal(t, expected, result)

	// Test case 4: Empty filters
	filters = ""
	fieldMap = map[string]string{
		"Id": "id",
	}
	expected = ""
	result = BuildQuery(filters, fieldMap)
	require.Equal(t, expected, result)
}

func TestGetSQLOperator(t *testing.T) {
	// Test case 1: Equal operator
	symbol, value := getSQLOperator("eq", "1")
	require.Equal(t, "=", symbol)
	require.Equal(t, "1", value)

	// Test case 2: Less than operator
	symbol, value = getSQLOperator("lt", "100")
	require.Equal(t, "<", symbol)
	require.Equal(t, "100", value)

	// Test case 3: Greater than operator
	symbol, value = getSQLOperator("gt", "100")
	require.Equal(t, ">", symbol)
	require.Equal(t, "100", value)

	// Test case 4: Not equal operator
	symbol, value = getSQLOperator("neq", "1")
	require.Equal(t, "!=", symbol)
	require.Equal(t, "1", value)

	// Test case 5: Between operator
	symbol, value = getSQLOperator("btwn", "1 AND 10")
	require.Equal(t, "BETWEEN", symbol)
	require.Equal(t, "1 AND 10", value)

	// Test case 6: Starts with operator
	symbol, value = getSQLOperator("sw", "test")
	require.Equal(t, "ILIKE ", symbol)
	require.Equal(t, "test%", value)

	// Test case 7: Ends with operator
	symbol, value = getSQLOperator("ew", "test")
	require.Equal(t, "ILIKE ", symbol)
	require.Equal(t, "%test", value)

	// Test case 8: Contains operator
	symbol, value = getSQLOperator("has", "test")
	require.Equal(t, "ILIKE ", symbol)
	require.Equal(t, "%test%", value)

	// Test case 9: Invalid operator
	symbol, value = getSQLOperator("invalid", "test")
	require.Equal(t, "-", symbol)
	require.Equal(t, "test", value)
}
