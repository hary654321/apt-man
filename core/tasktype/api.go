package tasktype

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
	"zrDispatch/common/log"
	"zrDispatch/core/utils/resp"
)

var _ TaskRuner = DataAPI{}

// DataAPI http req task
type DataAPI struct {
	URL     string            `json:"url" comment:"URL"`
	Method  string            `json:"method" comment:"Method"`
	PayLoad string            `json:"payload" comment:"PayLoad"`
	Header  map[string]string `json:"header" comment:"Header"`
}

// Header
// Body
// Test

// Type return api
func (da DataAPI) Type() string {
	return "api"
}

// Run implment TaskRun interface
func (da DataAPI) Run(ctx context.Context) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		var exitCode = DefaultExitCode
		defer pw.Close()
		defer func() {
			now := time.Now().Local().Format("2006-01-02 15:04:05: ")
			pw.Write([]byte(fmt.Sprintf("\n%sRun Finished,Return Code:%5d", now, exitCode))) // write exitCode,total 5 byte
			// pw.Write([]byte(fmt.Sprintf("%3d", exitCode))) // write exitCode,total 3 byte
		}()
		// go1.13 use NewRequestWithContext

		req, err := http.NewRequestWithContext(ctx, da.Method, da.URL, bytes.NewReader([]byte(da.PayLoad)))
		if err != nil {
			pw.Write([]byte(err.Error()))
			log.Error("NewRequest failed", zap.Error(err))
			return
		}

		for k, v := range da.Header {
			req.Header.Add(k, v)
		}

		client := http.DefaultClient
		doresp, err := client.Do(req)
		if err != nil {
			log.Error("client Do failed", zap.Error(err))
			var customerr bytes.Buffer
			switch ctx.Err() {
			case context.DeadlineExceeded:
				customerr.WriteString(resp.GetMsg(resp.ErrCtxDeadlineExceeded))
			case context.Canceled:
				customerr.WriteString(resp.GetMsg(resp.ErrCtxCanceled))
			default:
				customerr.WriteString(err.Error())
			}
			pw.Write(customerr.Bytes())
			return
		}
		defer doresp.Body.Close()

		bs, err := ioutil.ReadAll(doresp.Body)
		if err != nil {
			log.Error("Read failed", zap.Error(err))
			return
		}
		pw.Write(bs)

		if doresp.StatusCode > 0 {
			exitCode = doresp.StatusCode
		}
	}()
	return pr
}
