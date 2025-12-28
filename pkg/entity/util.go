package entity

import (
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/thoas/go-funk"
)

var excludedFields map[string]struct{}

func init() {

	set := funk.ToSet([]string{"CreatedById", "UpdatedById", "UpdatedAt", "CreatedAt"})
	excludedFields = set.(map[string]struct{})
}

func GetEntityAttributes(input interface{}) (string, []interface{}) {
	result := make([]interface{}, 0)
	t := reflect.TypeOf(input)

	inputName := t.Name()
	// Dereference pointers to get the underlying type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Only proceed if we're dealing with a struct
	if t.Kind() != reflect.Struct {
		return "", result
	}

	// Iterate through all fields of the struct type
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if _, ok := excludedFields[strcase.ToCamel(field.Name)]; ok {
			continue
		}

		// Handle embedded/anonymous structs or pointers to structs
		if field.Anonymous {
			var embeddedType reflect.Type

			if field.Type.Kind() == reflect.Ptr {
				embeddedType = field.Type.Elem()
			} else {
				embeddedType = field.Type
			}

			// Make sure it's a struct before recursing
			if embeddedType.Kind() == reflect.Struct {
				embeddedInstance := reflect.New(embeddedType).Interface()
				_, embeddedFields := GetEntityAttributes(embeddedInstance)
				result = append(result, embeddedFields...)
			}
		} else {

			//If its a reference struct , then skip it
			if isReferenceOrChild(field) {
				continue
			}

			// Add the current field name to the result
			fieldName := strcase.ToCamel(field.Name)
			result = append(result, fieldName)
		}

	}

	return inputName, result
}

func isReferenceOrChild(field reflect.StructField) bool {

	var embeddedType reflect.Type
	if field.Type.Kind() == reflect.Ptr {
		embeddedType = field.Type.Elem()
	} else {
		embeddedType = field.Type
	}

	return embeddedType.Kind() == reflect.Struct || embeddedType.Kind() == reflect.Array || embeddedType.Kind() == reflect.Slice
}
