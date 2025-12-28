package util

import (
	"bytes"

	"github.com/xuri/excelize/v2"
)

func ReadData(data []byte) *excelize.File {
	exlz, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil
	}

	return exlz
}
