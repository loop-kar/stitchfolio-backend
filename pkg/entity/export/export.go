package entity

import (
	"fmt"
	"maps"

	"github.com/imkarthi24/sf-backend/pkg/entity"
	"github.com/imkarthi24/sf-backend/pkg/util"
	"github.com/xuri/excelize/v2"
)

func ExportEntityData(entityData map[string][]interface{}) *excelize.File {

	excel := excelize.NewFile()

	defer excel.Close()

	//given an interface , get all the attributes
	for entityName, values := range entityData {

		//no records to process , then continue with next
		if len(values) == 0 {
			continue
		}

		_, attributes := entity.GetEntityAttributes(values[0])

		//create sheet with entity name
		index, err := excel.NewSheet(entityName)
		if err != nil {
			return nil
		}
		excel.SetActiveSheet(index)

		//insert headers in first row
		excel.SetSheetRow(entityName, headerRowStart, &attributes)

		for i, ent := range values {
			entityMap := util.ToTypeMap(ent)
			excel.SetSheetRow(entityName, fmt.Sprintf(dataRowStartFormat, i+65), maps.Values(entityMap))
		}
	}

	excel.DeleteSheet("Sheet1")
	excel.SetActiveSheet(0)

	if err := excel.SaveAs("EntityExport.xlsx"); err != nil {
		fmt.Println(err)
	}

	return excel
}
