package util

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func BuildQuery(filters string, fieldMap map[string]string) string {

	filterArray := strings.Split(filters, ",")

	queryString := ""
	for _, filter := range filterArray {

		parts := strings.Split(filter, " ")

		// id eq 1 , amount eq 1000 , name eq 'SWE'
		//field is usually in small camelCase  eg : camelCase not CamelCase
		field, symbol, value := parts[0], parts[1], parts[2]

		dbField := fieldMap[strcase.ToCamel(field)]

		// given filter does not exist for this table
		if IsNilOrEmptyString(&dbField) {
			continue
		}

		if !IsNilOrEmptyString(&queryString) {
			queryString = queryString + " AND "
		}

		sqlSymbol, value := getSQLOperator(symbol, value)

		queryString += fmt.Sprintf("%s %s %s", dbField, sqlSymbol, value)
	}

	return queryString

}

func getSQLOperator(symbol, value string) (string, string) {
	switch symbol {
	case "eq":
		return "=", value
	case "lt":
		return "<", value
	case "gt":
		return ">", value
	case "neq":
		return "!=", value
	case "btwn":
		return "BETWEEN", value
	case "sw":
		return "ILIKE ", InsertOperatorAtPosition(value, "%", -2)
	case "ew":
		return "ILIKE ", InsertOperatorAtPosition(value, "%", 1)
	case "has":
		value = EncloseWithPercentageOperator(value)
		return "ILIKE ", value

	}
	return "-", value
}
