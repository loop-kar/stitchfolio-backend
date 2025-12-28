package entity

import (
	"encoding/json"

	"github.com/iancoleman/strcase"
	"github.com/imkarthi24/sf-backend/pkg/util"
)

// GetExcelData return a map[string][]string, where key is the entity name
// and values are rows in excel in json format along with headers
// eg:{User : []string{"json of user1","json of user2"}}

func GetExcelData(data []byte) map[string][]string {
	//convert to excel
	excel := util.ReadData(data)

	defer func() {
		err := excel.Close()
		if err != nil {
			return
		}
	}()

	//iterate thru each sheet to see what entities we have
	sheets := excel.GetSheetList()
	dataMap := make(map[string][]string, len(sheets))
	for _, sheet := range sheets {

		rows, err := excel.GetRows(sheet)
		if err != nil {
			return nil
		}

		jsonRowsData := processRows(rows)
		dataMap[sheet] = jsonRowsData
	}

	return dataMap
}

func processRows(rows [][]string) []string {

	result := make([]string, 0, len(rows)-1)

	headers := rows[0]

	//1st row is headers , skip and start from next
	for _, row := range rows[1:] {
		rowMap := make(map[string]string)

		for i, cell := range row {

			if i < len(headers) {
				rowMap[strcase.ToLowerCamel(headers[i])] = cell
			}
		}

		jsonBytes, err := json.Marshal(rowMap)
		if err != nil {
			// Handle error - you might want to customize this based on your error handling strategy
			continue
		}

		// Add the JSON string to result
		result = append(result, string(jsonBytes))
	}

	return result
}
