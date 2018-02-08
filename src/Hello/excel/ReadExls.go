package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)
//go get github.com/tealeg/xlsx
func main() {
	excelFileName := "MyXLSXFile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Print(nil)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}