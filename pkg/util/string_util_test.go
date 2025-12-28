package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNilOrEmptyString(t *testing.T) {

	inputString := ""
	actualValue := IsNilOrEmptyString(&inputString)
	expectedvalue := true

	require.Equal(t, expectedvalue, actualValue)

	inputString1 := "testing"
	actualValue1 := IsNilOrEmptyString(&inputString1)
	expectedvalue1 := false

	require.Equal(t, expectedvalue1, actualValue1)

}

func TestIsNilOrEmptyString_Negative(t *testing.T) {

	inputString := "NotNull"
	actualValue := IsNilOrEmptyString(&inputString)
	expectedvalue := true

	require.NotEqual(t, expectedvalue, actualValue)

	inputString1 := ""
	actualValue1 := IsNilOrEmptyString(&inputString1)
	expectedvalue1 := false

	require.NotEqual(t, expectedvalue1, actualValue1)

}

func TestCSVToArray(t *testing.T) {
	inputString := ""
	actualValue, err := CSVToArray(&inputString, ",")
	expectedvalue := []int32([]int32(nil))

	require.Nil(t, err)
	require.Equal(t, expectedvalue, actualValue)

	inputString1 := "1,2,3"
	actualValue1, err1 := CSVToArray(&inputString1, ",")
	expectedvalue1 := []int32{1, 2, 3}

	require.Nil(t, err1)
	require.Equal(t, expectedvalue1, actualValue1)

}

func TestCSVToArray_Negative(t *testing.T) {
	inputString := "1"
	actualValue, err := CSVToArray(&inputString, ",")
	expectedvalue := []int32([]int32(nil))

	require.Nil(t, err)
	require.NotEqual(t, expectedvalue, actualValue)

	inputString1 := ""
	actualValue1, err1 := CSVToArray(&inputString1, ",")
	expectedvalue1 := []int32{1, 2, 3}

	require.Nil(t, err1)
	require.NotEqual(t, expectedvalue1, actualValue1)

}

func TestArrayToCSV(t *testing.T) {

	inputString := []string{"this", "is", "pranesh"}
	actualValue, err := ArrayToCSV(inputString, ",")
	expectedvalue := "this,is,pranesh"

	require.Nil(t, err)
	require.Equal(t, expectedvalue, actualValue)

	inputString1 := []string{"1", "2", "3"}
	actualValue1, err1 := ArrayToCSV(inputString1, ",")
	expectedvalue1 := "1,2,3"

	require.Nil(t, err1)
	require.Equal(t, expectedvalue1, actualValue1)

	inputString2 := []string{}
	actualValue2, err2 := ArrayToCSV(inputString2, ",")
	expectedvalue2 := ""

	require.Nil(t, err2)
	require.Equal(t, expectedvalue2, actualValue2)

	inputString3 := []string{"this/is/tesing"}
	actualValue3, err3 := ArrayToCSV(inputString3, ",")
	expectedvalue3 := "this/is/tesing"

	require.Nil(t, err3)
	require.Equal(t, expectedvalue3, actualValue3)

}

func TestArrayToCSV_Negative(t *testing.T) {

	inputString := []string{"negativeTesting"}
	actualValue, err := ArrayToCSV(inputString, ",")
	expectedvalue := "negative,Testing"

	require.Nil(t, err)
	require.NotEqual(t, expectedvalue, actualValue)

	inputString1 := []string{"1", "2", "3"}
	actualValue1, err1 := ArrayToCSV(inputString1, ",")
	expectedvalue1 := ""

	require.Nil(t, err1)
	require.NotEqual(t, expectedvalue1, actualValue1)

	inputString2 := []string{}
	actualValue2, err2 := ArrayToCSV(inputString2, ",")
	expectedvalue2 := "empty"

	require.Nil(t, err2)
	require.NotEqual(t, expectedvalue2, actualValue2)

}
