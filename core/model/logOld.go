package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"zrDispatch/common/db"
	"zrDispatch/common/log"
	"zrDispatch/common/utils"
	"zrDispatch/core/utils/define"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	// Email check email
	LOG_STATUS_FAIL int = -1
	// Name check name
	LOG_STATUS_SUCC int = 1
)

// save task reps log
func SaveLog(ctx context.Context, l *define.Log) error {
	log.Info("start save tasklog", zap.String("task", l.Name))
	savesql := `INSERT INTO log
				(name,
				taskid,
				runTaskId,
				hostid,
				starttime,
				endtime,
				totalruntime,
				status,
				taskresps,
				triggertype,
				errcode,
				errmsg,
				errtasktype,
				errtaskid,
				errtask
			)
			VALUES
			(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, savesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	taskresps, err := json.Marshal(l.Taskresps)
	if err != nil {
		return fmt.Errorf("json.Marshal failed: %w", err)
	}
	log.Info("RunTaskID" + l.RunTaskID)
	_, err = stmt.ExecContext(ctx, l.Name, l.TaskID, l.RunTaskID, l.HostId,
		l.StartTime, l.EndTime, l.TotalRunTime,
		l.Status, taskresps, l.Trigger, l.ErrCode, l.ErrMsg,
		l.ErrTasktype, l.ErrTaskID, l.ErrTask)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	return nil
}

// update host last recv hearbeat time
func UpdateRes(ctx context.Context, runTaskId, progress, taskresps string) error {
	log.Info("UpdateRes", zap.String(runTaskId, taskresps))
	updatesql := `UPDATE log set progress=?,taskresps=? WHERE runTaskId=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, updatesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, progress, taskresps, runTaskId)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	line, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected failed: %w", err)
	}
	if line <= 0 {
		log.Error("修改失败")
	}
	return nil
}

func UpdateLogStatus(ctx context.Context, runTaskId string, status int) error {
	updatesql := `UPDATE log set status=? WHERE runTaskId=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, updatesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, status, runTaskId)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	line, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected failed: %w", err)
	}
	if line <= 0 {
		log.Error("修改失败")
	}
	return nil
}

// GetLog get task resp log by taskid
func GetLog(ctx context.Context, taskname string, status int, offset, limit int) ([]*define.Log, int, error) {
	logs := []*define.Log{}
	getsql := `SELECT 
					name,
					taskid,
					runTaskId,
					starttime,
					endtime,
					totalruntime,
					status,
					triggertype,
					errcode,
					errmsg,
					errtasktype,
					errtaskid,
					errtask,
					progress,
					taskresps,
					hostid
				FROM 
					log`
	args := []interface{}{}
	if taskname != "" {
		args = append(args, taskname+"%")
		getsql += ` WHERE name LIKE ?`
	}

	if status != 0 {
		if len(args) == 1 {
			getsql += ` AND status=?`
		} else {
			getsql += ` WHERE status=?`
		}

		args = append(args, status)
	}
	count, err := countColums(ctx, getsql, args...)
	if err != nil {
		return logs, 0, fmt.Errorf("countColums failed: %w", err)
	}
	getsql += ` ORDER BY id DESC LIMIT ? OFFSET ?`
	args = append(args, limit, offset)

	conn, err := db.GetConn(ctx)
	if err != nil {
		return logs, 0, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, getsql)
	if err != nil {
		return logs, 0, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return logs, 0, fmt.Errorf("stmt.QueryContext failed: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		getlog := define.Log{}
		err = rows.Scan(
			&getlog.Name,
			&getlog.TaskID,
			&getlog.RunTaskID,
			&getlog.StartTime,
			&getlog.EndTime,
			&getlog.TotalRunTime,
			&getlog.Status,
			&getlog.Trigger,
			&getlog.ErrCode,
			&getlog.ErrMsg,
			&getlog.ErrTasktype,
			&getlog.ErrTaskID,
			&getlog.ErrTask,
			&getlog.Progress,
			&getlog.Taskresps,
			&getlog.HostId,
		)
		if err != nil {
			log.Error("rows.Scan failed", zap.Error(err))
			continue
		}
		getlog.ErrTaskTypeStr = getlog.ErrTasktype.String()
		getlog.StartTimeStr = getlog.StartTime
		getlog.EndTimeStr = getlog.EndTime
		getlog.Triggerstr = getlog.Trigger.String()
		logs = append(logs, &getlog)
	}
	return logs, count, nil
}

// GetTreeLog get tree log data
func GetTreeLog(ctx context.Context, id string, startTime int64) ([]*define.TaskStatusTree, error) {
	runTaskId := id + utils.GetInterfaceToString(startTime)
	sqlget := `SELECT taskresps FROM log WHERE runTaskId=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, sqlget)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var taskreposbyte []byte
	err = stmt.QueryRowContext(ctx, runTaskId).Scan(&taskreposbyte)
	defer stmt.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			return make([]*define.TaskStatusTree, 0), nil
		}
		return nil, err
	}
	taskrepos := []*define.TaskResp{}

	log.Info(utils.GetInterfaceToString(taskreposbyte))
	if utils.GetInterfaceToString(taskreposbyte) == "null" {
		return make([]*define.TaskStatusTree, 0), nil
	}

	err = json.Unmarshal(taskreposbyte, &taskrepos)
	if err != nil {
		log.Error("model.GetTreeLog", zap.Error(err))
		return nil, err
	}
	retTasksStatus := define.GetTasksTreeStatus()
	task, err := GetTaskByID(ctx, id)
	switch err.(type) {
	case nil:
		goto Next
	case define.ErrNotExist:
		return retTasksStatus, nil
	default:
		return nil, err
	}
Next:
	// TODO 优化 只循环taskresps就可以取出
	if len(task.ParentTaskIds) != 0 {
		var isSet bool
		for _, taskid := range task.ParentTaskIds {
			var taskresp *define.TaskResp
			for _, task := range taskrepos {
				if taskid == task.TaskID && task.TaskType == define.ParentTask {
					taskresp = task
					break
				}
			}
			if taskresp == nil {
				continue
			}

			tasktreestatus := define.TaskStatusTree{
				Status:       taskresp.Status,
				ID:           taskid,
				Name:         taskresp.Task,
				TaskType:     define.ParentTask,
				TaskRespData: taskresp.LogData,
			}
			retTasksStatus[0].Children = append(retTasksStatus[0].Children, &tasktreestatus)

			if !isSet {
				// 如果存在fail那么节点的状态就是fail
				if taskresp.Status == define.TsFail.String() {
					retTasksStatus[0].Status = taskresp.Status
					isSet = true
				} else {
					retTasksStatus[0].Status = taskresp.Status
				}
			}
		}
		retTasksStatus[0].TaskType = define.ParentTask
	}

	var taskresp *define.TaskResp
	for _, task := range taskrepos {
		if id == task.TaskID && task.TaskType == define.MasterTask {
			taskresp = task
			break
		}
	}
	retTasksStatus[1].ID = taskresp.TaskID
	retTasksStatus[1].Name = taskresp.Task
	retTasksStatus[1].Status = taskresp.Status
	retTasksStatus[1].TaskRespData = taskresp.LogData
	retTasksStatus[1].TaskType = define.MasterTask

	if len(task.ChildTaskIds) != 0 {
		var isSet bool
		for _, id := range task.ChildTaskIds {
			var taskresp *define.TaskResp
			for _, task := range taskrepos {
				if id == task.TaskID && task.TaskType == define.ChildTask {
					taskresp = task
					break
				}
			}
			if taskresp == nil {
				continue
			}

			tasktreestatus := define.TaskStatusTree{
				Status:       taskresp.Status,
				ID:           id,
				Name:         taskresp.Task,
				TaskType:     define.ParentTask,
				TaskRespData: taskresp.LogData,
			}
			retTasksStatus[2].Children = append(retTasksStatus[2].Children, &tasktreestatus)

			if !isSet {
				// 如果存在fail那么节点的状态就是fail
				if taskresp.Status == define.TsFail.String() {
					retTasksStatus[2].Status = taskresp.Status
					isSet = true
				} else {
					retTasksStatus[2].Status = taskresp.Status
				}
			}
		}
		retTasksStatus[2].TaskType = define.ChildTask
	}
	return retTasksStatus, nil
}

// CleanTaskLog clean old task from time ago
func CleanTaskLog(ctx context.Context, name, taskid string, deletetime int64) (int64, error) {
	delsql := `DELETE FROM log WHERE starttime < ?`
	args := []interface{}{deletetime}
	if name != "" {
		delsql += " AND name=? "
		args = append(args, name)
	} else if taskid != "" {
		delsql += " AND taskid=? "
		args = append(args, taskid)
	} else {
		return 0, errors.New("no delete id or name")
	}
	conn, err := db.GetConn(ctx)
	if err != nil {
		return 0, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, delsql)
	if err != nil {
		return 0, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return 0, fmt.Errorf("stmt.ExecContext failed: %w", err)
	}

	delcount, _ := res.RowsAffected()

	return delcount, nil
}

// SaveOperateLog save all user change operate
func SaveOperateLog(ctx context.Context,
	c *gin.Context, uid, username string,
	role define.Role, method, module, modulename string,
	operatetime int64, desc string, columns []define.Column) error {
	if c.GetInt("statuscode") != 0 {
		log.Error("req status code is not 0, do not save", zap.Int("statuscode", c.GetInt("statuscode")))
		return errors.New("return code is not equal 0")
	}
	log.Debug("start save operate", zap.String("username", username))
	operatesql := `INSERT INTO operate
			(uid,
			username,
			role,
			method,
			module,
			modulename,
			operatetime,
			description,
			columns)
			VALUES
			(
				?,?,?,?,?,?,?,?,?
			)
		`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, operatesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	columnsdata, err := json.Marshal(columns)
	_, err = stmt.ExecContext(ctx, uid, username, role, method, module, modulename, operatetime, desc, columnsdata)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	return nil
}

// GetOperate get operate log
func GetOperate(ctx context.Context, uid, username, method, module string, limit, offset int) ([]define.OperateLog, int, error) {
	getsql := `SELECT 
					uid,username,role,method,module,modulename, operatetime,description,columns
			   FROM 
					operate`
	query := []string{}
	args := []interface{}{}
	var count int

	if uid != "" {
		query = append(query, " uid=? ")
		args = append(args, uid)
	}
	if username != "" {
		query = append(query, " username=? ")
		args = append(args, username)
	}
	if method != "" {
		query = append(query, " method=? ")
		args = append(args, method)
	}
	if module != "" {
		query = append(query, " module=? ")
		args = append(args, module)
	}

	if len(query) > 0 {
		getsql += "WHERE"
		getsql += strings.Join(query, "AND")
	}
	oplogs := make([]define.OperateLog, 0, limit)

	if limit > 0 {
		var err error
		count, err = countColums(ctx, getsql, args...)
		if err != nil {
			return oplogs, 0, fmt.Errorf("countColums failed: %w", err)
		}
		getsql += ` ORDER BY id DESC LIMIT ? OFFSET ?`
		args = append(args, limit, offset)
	}

	conn, err := db.GetConn(ctx)
	if err != nil {
		return oplogs, 0, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, getsql)
	if err != nil {
		return oplogs, 0, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return oplogs, 0, fmt.Errorf("stmt.QueryContext failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			err         error
			columnsdata []byte
			oplog       define.OperateLog
			operatetime int64
		)
		// uid,username,role,method,module,modulename, operatetime,columns
		err = rows.Scan(&oplog.UID,
			&oplog.UserName,
			&oplog.Role,
			&oplog.Method,
			&oplog.Module,
			&oplog.ModuleName,
			&operatetime,
			&oplog.Desc,
			&columnsdata,
		)
		if err != nil {
			log.Error("rows.Scan failed", zap.Error(err))
			continue
		}

		oplog.OperateTime = utils.UnixToStr(operatetime)
		var columns []define.Column
		err = json.Unmarshal(columnsdata, &columns)
		if err != nil {
			log.Error("json.Unmarshal failed", zap.Error(err))
			continue
		}
		oplog.Columns = columns

		oplogs = append(oplogs, oplog)
	}
	return oplogs, count, nil
}

// SaveNewNotify save new notify
func SaveNewNotify(ctx context.Context, notify define.Notify) error {
	savesql := `INSERT INTO notify
				(
					notyfytype,
					notifyuid,
					title,
					content,
					is_read,
					notifytime
				)
			  VALUES 
				(
					?,?,?,?,?,?
				)
					`
	log.Debug("start save new notify", zap.Any("notify", notify))
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, savesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx,
		notify.NotifyType,
		notify.NotifyUID,
		notify.Title,
		notify.Content,
		false,
		notify.NotifyTime,
	)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}

	return nil
}

// GetNotifyByUID get user's notify
func GetNotifyByUID(ctx context.Context, uid string) ([]define.Notify, error) {
	getsql := `SELECT 
					id,notyfytype,title,content,notifytime
			   FROM 
					notify 
			   WHERE 
			   		is_read=? AND notifyuid=?`
	notifys := []define.Notify{}
	conn, err := db.GetConn(ctx)
	if err != nil {
		return notifys, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, getsql)
	if err != nil {
		return notifys, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, false, uid)
	if err != nil {
		return notifys, fmt.Errorf("stmt.QueryContext failed: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var notify define.Notify
		err = rows.Scan(&notify.ID, &notify.NotifyType, &notify.Title, &notify.Content, &notify.NotifyTime)
		if err != nil {
			log.Error("rows.Scan failed", zap.Error(err))
			continue
		}
		notify.NotifyTypeDesc = notify.NotifyType.String()
		notify.NotifyTimeDesc = utils.UnixToStr(notify.NotifyTime)
		notifys = append(notifys, notify)
	}
	return notifys, nil
}

// NotifyRead make notify status is readed
func NotifyRead(ctx context.Context, id int, uid string) error {
	args := []interface{}{true, uid}
	updatesql := `UPDATE notify SET is_read=? WHERE notifyuid=?`
	if id != 0 {
		updatesql += "  AND id=?"
		args = append(args, id)
	}
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, updatesql)
	defer stmt.Close()
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	return nil
}
