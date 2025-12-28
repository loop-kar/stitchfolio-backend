package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func IsNilOrEmptyString(s *string) bool {
	return s == nil || *s == ""
}

func CSVToArray(s *string, seperator string) ([]int32, error) {
	if IsNilOrEmptyString(s) {
		return nil, nil
	}

	result := make([]int32, 0)
	stringArr := strings.Split(*s, seperator)
	funk.ForEach(stringArr, func(x string) {
		value, _ := strconv.Atoi(x)
		result = append(result, int32(value))
	})
	return result, nil
}

func ArrayToCSV(s []string, seperator string) (string, error) {

	result := ""
	for _, a := range s {
		result = result + a + seperator
	}

	return strings.Trim(result, seperator), nil
}

func InsertOperatorAtPosition(str, operator string, position int) string {
	if position < 0 {
		position = len(str) + position + 1
	}
	return fmt.Sprintf("%s%s%s", str[:position], operator, str[position:])
}

func EncloseWithSymbol(str, symbol string) string {
	if IsNilOrEmptyString(&str) {
		return str
	}
	return fmt.Sprintf("%s%s%s", symbol, str, symbol)
}

func EncloseWithSingleQuote(str string) string {
	return EncloseWithSymbol(str, "'")
}

func EncloseWithPercentageOperator(str string) string {
	str = InsertOperatorAtPosition(str, "%", -2)
	str = InsertOperatorAtPosition(str, "%", 1)
	return str
}

func SplitNonEmpty(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}
