package utils

import (
	"bytes"
	"encoding/json"
	"os"
	"time"
	"zrDispatch/core/slog"
)

func WriteJson(path string, data interface{}) {

	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)

	err := enc.Encode(data)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_APPEND, 0764)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	jsonBuf := append(buf.Bytes(), []byte("\r\n")...)

	slog.Println(slog.DEBUG, jsonBuf)

	f.Write(buf.Bytes())

}

func WriteJsonLog(data interface{}) {

	path := time.Now().Format("2006-01-02") + ".log"
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)

	err := enc.Encode(data)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_APPEND, 0764)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	jsonBuf := append(buf.Bytes(), []byte("\r\n")...)

	slog.Println(slog.DEBUG, string(jsonBuf))

	f.Write(buf.Bytes())

}
