package utils

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

const (
	DEFAULT_SHEET_NAME = "Sheet1"
)

func ExportExcel(headName []string, rows [][]interface{}, w io.Writer) error {
	head := make([]interface{}, len(headName))
	for i, v := range headName {
		head[i] = v
	}

	f := excelize.NewFile()
	// 设置单元格的值
	defer f.Close()
	sheet, err := f.NewSheet(DEFAULT_SHEET_NAME)
	if err != nil {
		return fmt.Errorf(" create sheet error: %v", err)
	}
	stream, err := f.NewStreamWriter(DEFAULT_SHEET_NAME)
	if err != nil {
		return fmt.Errorf(" create stream writer error: %v", err)
	}
	stream.SetRow("A1", head)
	for i, row := range rows {
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		stream.SetRow(cell, row)
	}
	stream.Flush()
	f.SetActiveSheet(sheet)

	if _, err = f.WriteTo(w); err != nil {
		return fmt.Errorf(" write to file error: %v", err)
	}
	return nil
}

func ReadExcelByString(data string, sheetName string) (fileData [][]string, err error) {
	xlsx, err := excelize.OpenReader(bytes.NewBufferString(data))
	if err != nil {
		return nil, fmt.Errorf(" open file error: %v", err)
	}
	rows, err := xlsx.GetRows(sheetName)
	return rows, fmt.Errorf(" read file error: %v", err)
}

func ReadExcelByfile(filePath string, sheetName string) ([][]string, error) {
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf(" open file error: %v", err)
	}
	rows, err := xlsx.GetRows(sheetName)
	return rows, fmt.Errorf(" read file error: %v", err)
}
