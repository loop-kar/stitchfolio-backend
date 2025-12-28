package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func TestReadData(t *testing.T) {
	// Test case 1: Valid Excel file
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "Test Data")

	// Save to buffer
	buffer, err := f.WriteToBuffer()
	require.Nil(t, err)

	// Read the data
	result := ReadData(buffer.Bytes())
	require.NotNil(t, result)

	// Verify content
	value, err := result.GetCellValue(sheet, "A1")
	require.Nil(t, err)
	require.Equal(t, "Test Data", value)

	// Test case 2: Invalid Excel file
	invalidData := []byte("not an excel file")
	result = ReadData(invalidData)
	require.Nil(t, result)

	// Test case 3: Empty data
	emptyData := []byte{}
	result = ReadData(emptyData)
	require.Nil(t, result)
}
