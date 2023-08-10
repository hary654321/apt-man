package utils

import (
	"strconv"
	"time"
)

var HouarArr = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}

// UnixToStr change int to strc
func UnixToStr(i int64) string {
	// return time.Unix(i,0).Location().String()
	return time.Unix(i, 0).Local().Format("2006-01-02 15:04:05")

}

// StrToUnix change str to unix
func StrToUnix(t string) int64 {
	tparse, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		tparse, err = time.Parse("2006-01-02T15:04:05Z", t)
		if err != nil {
			return 0
		}
	}
	return tparse.Unix()
}

func GetTime() string {
	time := time.Now().Unix()
	s := strconv.FormatInt(time, 10)
	return s
}

func GetHaoMiao() string {
	time := time.Now().UnixMilli()
	s := GetInterfaceToString(time)
	return s
}

func GetTimeBefore(j int64) string {
	time := time.Now().Unix()
	time = time - j
	s := strconv.FormatInt(time, 10)
	return s
}
