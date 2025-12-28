package entity

import (
	"fmt"

	"github.com/imkarthi24/sf-backend/pkg/entity"
	"github.com/xuri/excelize/v2"
)

const headerRowStart string = "A1"
const dataRowStartFormat string = "%s1"

func ExportTemplate(entities []interface{}) *excelize.File {

	excel := excelize.NewFile()

	defer excel.Close()

	//given an interface , get all the attributes
	for _, e := range entities {
		entityName, attributes := entity.GetEntityAttributes(e)

		//create sheet with entity name
		index, err := excel.NewSheet(entityName)
		if err != nil {
			return nil
		}
		excel.SetActiveSheet(index)

		//insert headers in first row
		excel.SetSheetRow(entityName, headerRowStart, &attributes)
	}

	excel.DeleteSheet("Sheet1")
	excel.SetActiveSheet(0)

	if err := excel.SaveAs("EntityExport.xlsx"); err != nil {
		fmt.Println(err)
	}

	return excel
}
