package entity

import (
	"fmt"
	"os"
	"testing"
)

func Test_GetExcelData(t *testing.T) {

	data, _ := readContentFromFile("./../../../test/dump.xlsx")
	res := GetExcelData(data)
	fmt.Print(res)
}

func readContentFromFile(fileName string) ([]byte, error) {

	bs, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return bs, nil
}
