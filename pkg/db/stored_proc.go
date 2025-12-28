package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/imkarthi24/sf-backend/pkg/util"
	"gorm.io/gorm"
)

type StoredProcExecutor struct {
	db *gorm.DB
}

func (exe *StoredProcExecutor) CallStoredProcedure(ctx *context.Context, procName string, parameters map[string]interface{}) ([]ResultSet, error) {

	spResult := make([]ResultSet, 0)
	convertedQuery := exe.convertQuery(procName, parameters)

	result, err := exe.db.Raw(convertedQuery).Rows()
	if err != nil {
		return nil, err
	}

	resultSetAvailable := true

	for resultSetAvailable {

		rows, err := exe.prepareDataSet(result, false)
		if err != nil {
			return nil, err
		}

		res := ResultSet{
			Result: rows,
		}

		spResult = append(spResult, res)
		resultSetAvailable = result.NextResultSet()
	}
	result.Close()
	return spResult, nil
}

func (exe *StoredProcExecutor) convertQuery(procName string, parameters map[string]interface{}) string {

	procName = strings.TrimSpace(procName)
	query := fmt.Sprintf("SELECT * FROM %s", procName)
	seperator := ", "
	paramQuery := ""
	for k, v := range parameters {

		if v == nil {
			v = "NULL"
		} else if util.IsString(v) {
			v = util.EncloseWithSingleQuote(v.(string))
		}

		param := fmt.Sprintf("%s => %v %s", k, v, seperator)
		paramQuery = paramQuery + param
	}
	paramQuery = strings.Trim(paramQuery, seperator)

	return fmt.Sprintf("%s(%s)", query, paramQuery)

}

// Prepare the return dataset for select statements.
//
// Source: https://kylewbanks.com/blog/query-result-to-map-in-golang
func (exe *StoredProcExecutor) prepareDataSet(rows *sql.Rows, closeRows bool) ([]map[string]interface{}, error) {

	if closeRows {
		defer rows.Close()
	}

	var data []map[string]interface{}
	cols, _ := rows.Columns()

	// createDriver a slice of interface{}'s to represent each column
	// and a second slice to contain pointers to each item in the columns slice
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))

	for i := range columns {
		columnPointers[i] = &columns[i]
	}

	for rows.Next() {
		// scan the result into the column pointers
		err := rows.Scan(columnPointers...)
		if err != nil {
			return nil, err
		}

		// createDriver our map, and retrieve the value for each column from the pointers slice
		// storing it in the map with the name of the column as the key
		row := make(map[string]interface{})

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			row[colName] = *val
		}

		data = append(data, row)
	}

	return data, nil
}

type ResultSet struct {
	Result []map[string]interface{}
}
