package tasktype

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"zrDispatch/common/log"

	"go.uber.org/zap"

	pb "zrDispatch/core/proto"
	"zrDispatch/core/utils/define"
)

const (
	// DefaultExitCode default err code if not get run task code
	DefaultExitCode int = -1
)

// TaskRuner run task interface
// Please Implment io.ReadCloser
// reader last 3 byte must be exit code
type TaskRuner interface {
	Run(ctx context.Context) (out io.ReadCloser)
	Type() string
}

// get api or code  获取任务的信息  解json
func GetDataRun(t *pb.TaskReq) (TaskRuner, error) {
	switch define.TaskType(t.TaskType) {
	case define.TYPE_PORT:
		var code DataCode
		err := json.Unmarshal(t.TaskData, &code)
		if err != nil {
			return nil, err
		}
		code.LangDesc = code.Lang.String()
		return code, err

	default:
		err := fmt.Errorf("Unsupport TaskType %d", t.TaskType)
		return nil, err
	}
}

func GetTaskData(t *pb.TaskReq) define.TaskData {

	var taskParam define.TaskData
	err := json.Unmarshal(t.TaskData, &taskParam)
	if err != nil {
		log.Error("BindQuery offset failed", zap.Error(err))
	}

	return taskParam

}
