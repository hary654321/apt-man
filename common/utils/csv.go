package utils

import (
	"strings"
)

func GetcsvData(filename string) []map[string]any {

	rows, _ := ReadLineData(filename)
	head := rows[0]

	headarr := strings.Split(head, ",")

	data := rows[1:]
	var ResData []map[string]any

	for _, row := range data {
		var rowData = make(map[string]any)
		rowArr := strings.Split(row, ",")
		for k, v := range headarr {
			rowData[v] = rowArr[k]
			// slog.Println(slog.DEBUG, v, "========", rowArr[k])
		}
		ResData = append(ResData, rowData)
	}

	return ResData
}
