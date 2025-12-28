package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestExcelData struct {
	Name     string
	Age      int
	IsActive bool
}

func TestExcelInfo_BuildExcel(t *testing.T) {
	// Test case 1: Basic Excel creation with preset headers
	data := []TestExcelData{
		{Name: "John", Age: 30, IsActive: true},
		{Name: "Jane", Age: 25, IsActive: false},
	}

	excelInfo := &ExcelInfo{
		SheetName:        "TestSheet",
		Rows:             data,
		UsePresetHeaders: true,
	}

	file, err := excelInfo.BuildExcel()
	require.Nil(t, err)
	require.NotNil(t, file)

	// Verify headers
	value, err := file.GetCellValue("TestSheet", "A1")
	require.Nil(t, err)
	require.Equal(t, "Name", value)

	value, err = file.GetCellValue("TestSheet", "B1")
	require.Nil(t, err)
	require.Equal(t, "Age", value)

	value, err = file.GetCellValue("TestSheet", "C1")
	require.Nil(t, err)
	require.Equal(t, "IsActive", value)

	// Verify data
	value, err = file.GetCellValue("TestSheet", "A2")
	require.Nil(t, err)
	require.Equal(t, "John", value)

	value, err = file.GetCellValue("TestSheet", "B2")
	require.Nil(t, err)
	require.Equal(t, "30", value)

	value, err = file.GetCellValue("TestSheet", "C2")
	require.Nil(t, err)
	require.Equal(t, "true", value)

	// Test case 2: Custom headers
	excelInfo = &ExcelInfo{
		SheetName: "TestSheet",
		Headers:   []string{"Full Name", "Years", "Status"},
		Rows:      data,
	}

	file, err = excelInfo.BuildExcel()
	require.Nil(t, err)
	require.NotNil(t, file)

	// Verify custom headers
	value, err = file.GetCellValue("TestSheet", "A1")
	require.Nil(t, err)
	require.Equal(t, "Full Name", value)

	value, err = file.GetCellValue("TestSheet", "B1")
	require.Nil(t, err)
	require.Equal(t, "Years", value)

	value, err = file.GetCellValue("TestSheet", "C1")
	require.Nil(t, err)
	require.Equal(t, "Status", value)
}

func TestExcelInfo_Init(t *testing.T) {
	// Test case 1: Valid data
	data := []TestExcelData{
		{Name: "John", Age: 30, IsActive: true},
	}
	excelInfo := &ExcelInfo{
		SheetName: "TestSheet",
		Rows:      data,
	}
	err := excelInfo.init()
	require.Nil(t, err)

	// Test case 2: Invalid data (not a slice)
	excelInfo = &ExcelInfo{
		SheetName: "TestSheet",
		Rows:      "not a slice",
	}
	err = excelInfo.init()
	require.NotNil(t, err)
}

func TestExcelInfo_GetHeaders(t *testing.T) {
	data := []TestExcelData{
		{Name: "John", Age: 30, IsActive: true},
	}

	// Test case 1: Use preset headers
	excelInfo := &ExcelInfo{
		SheetName:        "TestSheet",
		Rows:             data,
		UsePresetHeaders: true,
	}
	err := excelInfo.init()
	require.Nil(t, err)
	headers := excelInfo.getHeaders()
	require.Equal(t, []string{"Name", "Age", "IsActive"}, headers)

	// Test case 2: Use custom headers
	excelInfo = &ExcelInfo{
		SheetName: "TestSheet",
		Headers:   []string{"Full Name", "Years", "Status"},
		Rows:      data,
	}
	headers = excelInfo.getHeaders()
	require.Equal(t, []string{"Full Name", "Years", "Status"}, headers)

	// Test case 3: Empty headers, use preset
	excelInfo = &ExcelInfo{
		SheetName: "TestSheet",
		Rows:      data,
	}
	err = excelInfo.init()
	require.Nil(t, err)
	headers = excelInfo.getHeaders()
	require.Equal(t, []string{"Name", "Age", "IsActive"}, headers)
}
