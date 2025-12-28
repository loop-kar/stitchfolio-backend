package util

import (
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/xuri/excelize/v2"
)

type ExcelInfo struct {
	SheetName        string
	Headers          []string
	Rows             interface{}
	UsePresetHeaders bool

	presetHeaders []string
}

var rows reflect.Value
var file *excelize.File

func (e *ExcelInfo) init() error {

	rows = reflect.ValueOf(e.Rows)
	e.fetchPresetHeaders()
	if rows.Kind() != reflect.Array && rows.Kind() != reflect.Slice {
		return fmt.Errorf("Invalid Data")
	}

	file = excelize.NewFile()

	return nil
}

func (e *ExcelInfo) BuildExcel() (*excelize.File, error) {

	err := e.init()
	if err != nil {
		return nil, err
	}

	defer func() error {
		err := file.Close()

		if err != nil {
			return err
		}

		return nil
	}()

	err = file.SetSheetName("Sheet1", e.SheetName)
	if err != nil {
		return nil, err
	}

	e.fillHeaders()
	e.fillData()
	e.addTable()

	return file, nil

}

func (e *ExcelInfo) fetchPresetHeaders() {

	value := rows.Index(0)

	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		normalizedFieldName := strcase.ToCamel(fieldName)
		e.presetHeaders = append(e.presetHeaders, normalizedFieldName)
	}
}

func (e *ExcelInfo) fillHeaders() error {

	for i, header := range e.getHeaders() {
		cell := fmt.Sprintf("%c1", 65+i)
		file.SetCellValue(e.SheetName, cell, header)
	}

	return nil
}

func (e *ExcelInfo) fillData() error {

	//Iterates through each item of the interface array
	startingRowNumber := 2

	for i := 0; i < rows.Len(); i++ {
		element := rows.Index(i)

		startingColumn := 65
		//Iterates through each element of the struct
		for i := 0; i < element.NumField(); i++ {
			field := element.Field(i)

			cell := fmt.Sprintf("%c%d", startingColumn, startingRowNumber)
			value := field.Interface()

			file.SetCellValue(e.SheetName, cell, value)
			startingColumn += 1

		}
		startingRowNumber += 1
	}

	return nil
}

func (e *ExcelInfo) addTable() error {

	cellRange := e.calculateRange()
	tableRange := fmt.Sprintf("%s:%s", cellRange[0], cellRange[1])
	err := file.AddTable(e.SheetName, &excelize.Table{
		Range:             tableRange,
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   true,
		ShowLastColumn:    true,
		ShowColumnStripes: true,
	})

	return err
}

func (e *ExcelInfo) calculateRange() []string {

	res := make([]string, 2)
	res[0] = "A1"
	columns := len((e.getHeaders())) - 1
	rows := rows.Len() + 1

	char := fmt.Sprintf("%c", 65+columns)

	res[1] = fmt.Sprintf("%s%d", char, rows)

	return res

}

func (e *ExcelInfo) getHeaders() []string {
	if e.UsePresetHeaders || len(e.Headers) == 0 {
		return e.presetHeaders
	}

	return e.Headers
}
