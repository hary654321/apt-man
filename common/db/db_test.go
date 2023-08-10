package db

import (
	"context"
	"os"
	"testing"
)

func TestNewDb(t *testing.T) {

	err := NewDb(Drivename("sqlite3"),
		Dsn("sqlite3.db"),
		MaxIdleConnection(10),
		MaxQueryTime(3),
		MaxQueryTime(3),
		MaxOpenConnection(3),
	)
	if err != nil {
		t.Fatalf("NewDb Err: %v", err)
	}
	conn, err := GetConn(context.Background())
	if err != nil {
		t.Fatalf("Get Conn Err: %v", err)
	}
	conn.Close()
	_ = os.Remove("sqlite3.db")
}

// 将指纹load进mysql  进行维护
//func InitFinger() {
//	fs, _ := fingerprintEmbed.Open(fingerprintPath)
//	sourceBuf, err := io.ReadAll(fs)
//	if err != nil {
//		slog.Printf(slog.ERROR, "%s", err.Error())
//	}
//	source := strings.Split(string(sourceBuf), "\n")
//	for _, line := range source {
//		line = strings.TrimSpace(line)
//		r := strings.SplitAfterN(line, "\t", 2)
//		if len(r) != 2 {
//			slog.Printf(slog.ERROR, "%s", err.Error())
//			continue
//		}
//		data := make(map[string]interface{})
//		data["name"] = r[0]
//		data["description"] = r[0]
//		data["finger"] = r[1]
//		models.AddFinger(data)
//
//	}
//}
