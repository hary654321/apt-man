package utils

import (
	"testing"
	"time"
)

func TestUnixToTime(t *testing.T) {
	var cstZone = time.FixedZone("CST", 7*3600) // 东八
	println(1111)
	println(time.Now().In(cstZone).Format("2006-01-02-15"))

}
