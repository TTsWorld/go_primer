package _struct

import (
	"fmt"
	"testing"
)

type ExcelSheetData struct {
	Title     string
	FileName  string
	SheetName string
	ColName   []interface{}
	RowData   []interface{}
}

var ExcelData ExcelSheetData

func TestName(t *testing.T) {
	fmt.Printf("%+v", ExcelData)

}
