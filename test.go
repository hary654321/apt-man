package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type NotUsed struct {
	Name string
}

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id            string  `csv:"client_id"`
	Name          string  `csv:"client_name"`
	Age           string  `csv:"client_age"`
	NotUsedString string  `csv:"-"`
	NotUsedStruct NotUsed `csv:"-"`
}

func main() {
	doc()
}

func doc() {

	tmpl, err := template.ParseFiles("a.doc")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	name := "cats"
	err = tmpl.Execute(os.Stdout, name)

}

func getPostArr(port string) (portRes []string) {
	if strings.Contains(port, ",") {
		portArr := strings.Split(port, ",")

		for _, v := range portArr {
			if strings.Contains(v, "-") {
				portRange := strings.Split(port, "-")
				fmt.Println("%#", portRange)
				startPort, _ := strconv.Atoi(portRange[0])
				endPort, _ := strconv.Atoi(portRange[1])

				for i := startPort; i <= endPort; i++ {

					portRes = append(portRes, strconv.Itoa(i))

				}
			} else {
				portRes = append(portRes, v)
			}
		}
	}

	return
}
