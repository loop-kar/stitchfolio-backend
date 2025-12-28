package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFindMonthsBetweenTwoDates(t *testing.T) {
	// Test case 1: Same month
	date1 := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	months := FindMonthsBetweenTwoDates(&date1, &date2)
	require.Equal(t, 1, months)

	// Test case 2: Different months
	date1 = time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	date2 = time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
	months = FindMonthsBetweenTwoDates(&date1, &date2)
	require.Equal(t, 2, months)

	// Test case 3: Date1 after Date2
	date1 = time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
	date2 = time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	months = FindMonthsBetweenTwoDates(&date1, &date2)
	require.Equal(t, -1, months)

	// Test case 4: Different years
	date1 = time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC)
	date2 = time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	months = FindMonthsBetweenTwoDates(&date1, &date2)
	require.Equal(t, 2, months)
}

func TestFindMonthsBetweenTwoDates_Date1AfterDate2(t *testing.T) {
	date1 := time.Now().Add(time.Hour - 48)
	date2 := time.Now()

	actualValue := FindMonthsBetweenTwoDates(&date2, &date1)
	expectedvalue := -1

	require.NotEqual(t, expectedvalue, actualValue)

}

func TestGenerateDateTimeFromYYYYMM(t *testing.T) {
	// Test case 1: Valid date string
	dateString := "2024-03"
	date := 15
	result, err := GenerateDateTimeFromYYYYMM(dateString, date)
	require.Nil(t, err)
	expected := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	require.Equal(t, expected, result)

	// Test case 2: Invalid date string format
	dateString = "2024/03"
	_, err = GenerateDateTimeFromYYYYMM(dateString, date)
	require.NotNil(t, err)

	// Test case 3: Invalid month
	dateString = "2024-13"
	_, err = GenerateDateTimeFromYYYYMM(dateString, date)
	require.NotNil(t, err)
}

func TestGenerateDateTimeFromYYYYMM_Negative(t *testing.T) {

	dateString := "2021-11"
	date := 12
	actualValue, err := GenerateDateTimeFromYYYYMM(dateString, date)
	expectedvalue, _ := time.Parse(time.DateOnly, "2021-10-12")

	require.Nil(t, err)
	require.NotEqual(t, expectedvalue, actualValue)

}

func TestTransformTimeToMMYYYY(t *testing.T) {
	// Test case 1: Single digit month
	date := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	result := TransformTimeToMMYYYY(&date)
	require.Equal(t, "03/2024", result)

	// Test case 2: Double digit month
	date = time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
	result = TransformTimeToMMYYYY(&date)
	require.Equal(t, "12/2024", result)
}

func TestTransformTimeToMMYYYY_Negative(t *testing.T) {

	date := time.Date(2021, 8, 15, 0, 0, 0, 0, time.Local)

	actualValue := TransformTimeToMMYYYY(&date)
	expectedvalue := "09/2021"
	require.NotEqual(t, expectedvalue, actualValue)

}

func TestConvertMToMM(t *testing.T) {
	// Test case 1: Single digit
	result := ConvertMToMM(5)
	require.Equal(t, "05", result)

	// Test case 2: Double digit
	result = ConvertMToMM(12)
	require.Equal(t, "12", result)

	// Test case 3: Zero
	result = ConvertMToMM(0)
	require.Equal(t, "00", result)
}

func TestConvertMToMM_Negative(t *testing.T) {
	date := 13
	actualValue := ConvertMToMM(date)
	expectedvalue := "12"
	require.NotEqual(t, expectedvalue, actualValue)

	date1 := -10
	actualValue1 := ConvertMToMM(date1)
	expectedvalue1 := "0"
	require.NotEqual(t, expectedvalue1, actualValue1)

}

func TestGenerateDateTimeFromString(t *testing.T) {
	// Test case 1: Valid date string
	dateString := "2024-03-15"
	result, err := GenerateDateTimeFromString(&dateString)
	require.Nil(t, err)
	expected := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	require.Equal(t, expected, *result)

	// Test case 2: Invalid date string
	dateString = "2024/03/15"
	_, err = GenerateDateTimeFromString(&dateString)
	require.NotNil(t, err)

	// Test case 3: Nil string
	var nilString *string
	result, err = GenerateDateTimeFromString(nilString)
	require.Nil(t, err)
	require.Nil(t, result)
}

func TestGenerateDateTimeFromString_Negative(t *testing.T) {

	dateSting := "2023-03-12"

	actualValue, err := GenerateDateTimeFromString(&dateSting)
	expectedvalue := time.Time(time.Date(2023, time.March, 13, 0, 0, 0, 0, time.UTC))

	require.Nil(t, err)
	require.NotEqual(t, expectedvalue, actualValue, " String is converted to Date time format ")

}

func TestToStringOrDefault(t *testing.T) {
	// Test case 1: Valid date
	date := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	result := DateTimeToStringOrDefault(&date, time.DateOnly)
	require.Equal(t, "2024-03-15", result)

	// Test case 2: Nil date
	result = DateTimeToStringOrDefault(nil, time.DateOnly)
	require.Equal(t, "", result)

	// Test case 3: Custom layout
	result = DateTimeToStringOrDefault(&date, "2006/01/02")
	require.Equal(t, "2024/03/15", result)
}

func TestIsEmptyDate(t *testing.T) {
	// Test case 1: Empty date
	emptyDate := time.Time{}
	require.True(t, IsEmptyDate(&emptyDate))

	// Test case 2: Non-empty date
	date := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	require.False(t, IsEmptyDate(&date))
}

func TestIsSamedate(t *testing.T) {
	// Test case 1: Same date
	date1 := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2024, 3, 15, 12, 30, 0, 0, time.UTC)
	require.True(t, IsSamedate(&date1, &date2))

	// Test case 2: Different dates
	date1 = time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	date2 = time.Date(2024, 3, 16, 0, 0, 0, 0, time.UTC)
	require.False(t, IsSamedate(&date1, &date2))

	// Test case 3: Nil dates
	require.False(t, IsSamedate(nil, nil))
	require.False(t, IsSamedate(&date1, nil))
	require.False(t, IsSamedate(nil, &date2))
}
