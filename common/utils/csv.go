package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"zrDispatch/core/slog"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
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

func GetcsvDataPro(filename string, mapa map[string]string) (ResData []map[string]any, err error) {

	file, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("文件打开失败: ", err)
		return
	}

	reader := csv.NewReader(transform.NewReader(bytes.NewReader(file), simplifiedchinese.GBK.NewDecoder()))
	rowNum := 1
	var headarr []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}

		if err != nil {
			fmt.Println("读取文件时发生错误: ", err)
			break
		}
		if rowNum == 1 {
			headarr = line
		} else {
			var rowData = make(map[string]any)
			for k, v := range headarr {

				if v == "" {
					continue
				}

				rowData[mapa[v]] = line[k]
				slog.Println(slog.DEBUG, v, "========", line[k])
			}

			ResData = append(ResData, rowData)
		}
		rowNum++

	}
	return
}
