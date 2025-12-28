package util

import (
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
)

type Type string

const (
	Integer   Type = "Integer"
	String    Type = "String"
	Boolean   Type = "Boolean"
	Float64   Type = "Float64"
	Interface Type = "Interface"
)

func GetType(val interface{}) Type {
	switch val.(type) {
	case int:
		return Integer
	case float64:
		return Float64
	case bool:
		return Boolean
	case string:
		return String
	default:
		return Interface
	}
}

func IsInteger(val interface{}) bool {
	_, ok := val.(int)
	return ok
}

func IsString(val interface{}) bool {
	_, ok := val.(string)
	return ok
}

func IsBool(val interface{}) bool {
	_, ok := val.(bool)
	return ok
}

func IsFloat64(val interface{}) bool {
	_, ok := val.(float64)
	return ok
}

func GetBooleanFields(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(input)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Bool {
			fieldName := field.Name
			fieldValue := v.Field(i).Bool()
			result[fieldName] = fieldValue
		}
	}

	return result
}

func PrepareEntityForUpdate(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(input)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	mapstructure.Decode(input, result)

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Anonymous {
			continue
		}

		value := v.Field(i).Interface()

		// Check if the field's value is not the zero value for its type
		if field.Type.Kind() == reflect.Bool || !reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
			fieldName := strcase.ToSnake(field.Name)
			result[fieldName] = value
		}
	}

	return result
}

func ToTypeMap(val interface{}) map[string]interface{} {
	res := make(map[string]interface{}, 0)
	mapstructure.Decode(val, res)
	return res
}

func IsEmptyValue(val interface{}) bool {
	return val == nil
}
