package utils

import (
	"strconv"
	"time"
)

func In_array(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func GenArr(len int) []int {
	var arr []int
	for i := 0; i < len; i++ {
		arr = append(arr, i)
	}

	return arr
}

func GetArrRand(arr []int) int {

	lent := len(arr)

	res := arr[RanNum(lent)]
	if res >= 0 {
		PrinfI("arr", arr)
		return res
	}
	if GetAliveCount(arr) < 1 {
		time.Sleep(10 * time.Second)
		return GetArrRand(arr)
	}
	return GetArrRand(arr)
}

func GetAliveCount(arr []int) int {
	i := 0
	for _, v := range arr {
		if v >= 0 {
			i++
		}
	}
	return i
}

func GetDieArrRand(arr []int) int {
	//PrinfI("arr", arr)
	lent := len(arr)

	k := RanNum(lent)
	res := arr[k]
	if res == -1 {
		return k
	}
	if GetDieCount(arr) < 1 {
		time.Sleep(1 * time.Second)
	}

	return GetDieArrRand(arr)
}

func GetDieCount(arr []int) int {
	i := 0
	for _, v := range arr {
		if v == -1 {
			i++
		}
	}

	return i
}

func StrArr2int(arrstr []string) (res []int) {
	for _, v := range arrstr {
		res = append(res, Str2int(v))
	}

	return
}

func Str2int(str string) int {
	a, _ := strconv.Atoi(str)

	return a
}

func IntArr2Str(arrstr []int) (res []string) {
	for _, v := range arrstr {
		res = append(res, GetInterfaceToString(v))
	}

	return
}
