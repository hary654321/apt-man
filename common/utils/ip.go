package utils

import (
	"bytes"
	"encoding/binary"
	"net"
	"regexp"
	"strconv"
	"strings"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

func GetIpUint32(str string) uint32 {
	netIP := net.ParseIP(str).To4()
	if netIP == nil {
		println(str)
	}

	u32 := binary.BigEndian.Uint32(netIP)

	return u32
}

func Long2IP(a uint32) string {
	ipInt := int(a)
	ipSegs := make([]string, 4)
	var len = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()

}

func GetAddrs(expr, ports string) (res []string) {
	ipArr := GetIpArr(expr)

	portArr := GetPortArr(ports)

	for _, i := range ipArr {
		for _, p := range portArr {
			slog.Println(slog.DEBUG, i+":"+p)
			res = append(res, i+":"+p)
		}
	}

	return
}

// 粒度问题会影响这里的设计
func GetIpArr(expr string) []string {
	var ipArr []string

	if strings.Contains(expr, ",") {
		ipArr = strings.Split(expr, ",")
	}

	if IsIPv4(expr) {
		ipArr = append(ipArr, expr)
	}

	if IsIPRanger(expr) {
		slog.Println(slog.DEBUG, expr)
		for _, ip := range RangerToIP(expr) {
			ipArr = append(ipArr, ip.String())
		}
	}

	if IsCIDR(expr) {
		slog.Println(slog.DEBUG, expr)
		for _, ip := range CIDRToIP(expr) {
			ipArr = append(ipArr, ip.String())
		}
	}

	return ipArr
}

func GetPortArr(port string) (portRes []string) {

	if strings.Contains(port, ",") {
		portArr := strings.Split(port, ",")

		for _, v := range portArr {
			if strings.Contains(v, "-") {
				portRange := strings.Split(port, "-")
				// fmt.Println("%#", portRange)
				startPort, _ := strconv.Atoi(portRange[0])
				endPort, _ := strconv.Atoi(portRange[1])

				for i := startPort; i <= endPort; i++ {

					portRes = append(portRes, strconv.Itoa(i))

				}
			} else {
				portRes = append(portRes, v)
			}
		}
	} else {
		if strings.Contains(port, "-") {
			portRange := strings.Split(port, "-")
			// fmt.Println("%#", portRange)
			startPort, _ := strconv.Atoi(portRange[0])
			endPort, _ := strconv.Atoi(portRange[1])

			for i := startPort; i <= endPort; i++ {

				portRes = append(portRes, strconv.Itoa(i))

			}
		}
	}

	if strings.Contains(port, "TOP") {
		num := port[3:]
		numInt, _ := strconv.Atoi(num)
		if numInt > 1000 {
			numInt = 1000
		}
		portRes = IntArr2Str(define.TOP_1000[0:numInt])
	}

	if isNumeric(port) {
		portRes = append(portRes, port)
	}

	//默认扫描前10个端口
	if len(portRes) == 0 {
		portRes = IntArr2Str(define.TOP_1000[0:10])
	}

	return
}

func isNumeric(input string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", input)
	return match
}
