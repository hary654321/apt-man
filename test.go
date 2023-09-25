package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
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
	println(time.Now().Unix())
	sleep()
	println(time.Now().Unix())
}

func sleep() {
	time.Sleep(time.Duration(1-1) * time.Second)
}

func encode() {
	res := base64.StdEncoding.EncodeToString([]byte("zw:1234qwer"))

	println(res)
}

func DecodeString() {
	decodedAuth, err := base64.StdEncoding.DecodeString("enc6MTIzNHF3ZXI=")
	if err != nil {
		// 处理解码错误
		return
	}
	// 按照冒号分割解码后的字符串，得到用户名和密码
	parts := bytes.Split(decodedAuth, []byte(":"))
	if len(parts) != 2 {
		// 处理格式错误
		return
	}
	username := string(parts[0])
	password := string(parts[1])

	println(username, password)
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
