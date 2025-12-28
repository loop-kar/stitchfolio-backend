package util

import (
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
)

func MapStructFieldsToCamelCase(e interface{}) map[string]string {
	dict := make(map[string]string)
	strct := reflect.Indirect(reflect.ValueOf(e))
	n := strct.NumField()
	for i := 0; i < n; i++ {
		fieldName := strct.Type().Field(i).Name
		camelStr := strcase.ToCamel(fieldName)
		dict[camelStr] = fieldName
	}
	return dict
}

func MapCamelStructFieldsToSnakeCase(e interface{}) map[string]string {
	dict := make(map[string]string)
	strct := reflect.Indirect(reflect.ValueOf(e))
	n := strct.NumField()
	for i := 0; i < n; i++ {
		fieldName := strct.Type().Field(i).Name
		camelStr := strcase.ToCamel(fieldName)
		dict[camelStr] = fieldName
	}
	return dict
}

func StructToMap(model interface{}) map[string]interface{} {

	responseMap := make(map[string]interface{})
	mapstructure.Decode(model, &responseMap)

	return responseMap
}
