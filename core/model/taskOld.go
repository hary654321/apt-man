package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"zrDispatch/common/db"
	"zrDispatch/common/log"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"

	"go.uber.org/zap"
)

// TASKTYPE
const (
	// Email check email
	RUN_STATUS_STOP int = iota
	// Name check name
	RUN_STATUS_RUNING
)

// change task
func ChangeTask(ctx context.Context, id string, run bool, tasktype int, taskData string,
	parentTaskIds []string, parentRunParallel bool, childTaskIds []string, childRunParallel bool,
	cronExpr string, timeout int, alarmUserIds []string, routePolicy define.RoutePolicy,
	alarmStatus define.AlarmStatus, hostGroupID, remark string) error {
	changesql := `UPDATE task 
					SET hostGroupID=?,
						run=?,
						taskType=?,
						taskData=?,
						parentTaskIds=?,
						parentRunParallel=?,
						childTaskIds=?,
						childRunParallel=?,
						cronExpr=?,
						timeout=?,
						alarmUserIds=?,
						routePolicy=?,
						alarmStatus=?,
						remark=?,
						updateTime=?
					WHERE id=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, changesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	updateTime := time.Now().Unix()

	_, err = stmt.ExecContext(ctx,
		hostGroupID,
		run,
		tasktype,
		taskData,
		strings.Join(parentTaskIds, ","),
		parentRunParallel,
		strings.Join(childTaskIds, ","),
		childRunParallel,
		cronExpr,
		timeout,
		strings.Join(alarmUserIds, ","),
		routePolicy,
		alarmStatus,
		remark,
		updateTime,
		id,
	)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	return nil
}

// DeleteTask delete task
func DeleteTask(ctx context.Context, id string) error {
	deletesql := `DELETE FROM task WHERE id=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, deletesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	return nil
}

// TaskIsUse check a task is other task's parent task ids or child task
func TaskIsUse(ctx context.Context, taskid string) (int, error) {
	querysql := `select count(*) from task WHERE id!=? AND (parentTaskIds LIKE ? OR childTaskIds LIKE ?) `
	conn, err := db.GetConn(ctx)
	if err != nil {
		return 0, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, querysql)
	if err != nil {
		return 0, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	var count int
	likequery := "%" + taskid + "%"
	err = stmt.QueryRowContext(ctx, taskid, likequery, likequery).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("stmt.QueryRowContext failed: %w", err)
	}
	return count, nil
}

// get all tasks
func GetTasks(ctx context.Context, offset, limit int, name string, runStatus int, createby string) ([]define.GetTask, int, error) {
	return getTasks(ctx, nil, name, offset, limit, true, runStatus, createby)
}

// GetTaskByID get task by id
func GetTaskByID(ctx context.Context, id string) (*define.GetTask, error) {
	tasks, _, err := getTasks(ctx, []string{id}, "", 0, 0, true, -1, "")
	if err != nil {
		return nil, err
	}
	if len(tasks) != 1 {
		err = define.ErrNotExist{Value: id}
		log.Error("get taskid failed", zap.Error(err))
		return nil, err
	}
	return &tasks[0], nil
}

// GetTaskByName get task by id
func GetTaskByName(ctx context.Context, name string) (*define.GetTask, error) {
	tasks, _, err := getTasks(ctx, nil, name, 0, 0, true, -1, "")
	if err != nil {
		return nil, err
	}
	if len(tasks) != 1 {
		err = define.ErrNotExist{Value: name}
		log.Error("get taskname failed", zap.Error(err))
		return nil, err
	}
	return &tasks[0], nil
}

// getTasks get takls by id
func getTasks(ctx context.Context,
	ids []string,
	name string,
	offset,
	limit int,
	first bool, /*Preventing endless loops*/
	runStatus int,
	createbyid string) ([]define.GetTask, int, error) {
	getsql := `SELECT t.id,
					t.name,
					t.group,
					t.tasktype,
					t.ip,
					t.port,
					t.probeId,
					t.plug,
					t.run,
					t.status,
					t.parentTaskIds,
					t.parentRunParallel,
					t.childTaskIds,
					t.childRunParallel,
					t.cronExpr,
					t.priority,
					t.timeout,
					t.threads,
					t.alarmUserIds,
					t.routePolicy,
					t.alarmStatus,
					u.name,
					t.createByID,
					hg.name,
					t.hostGroupID,
					t.remark,
					t.createTime,
					t.updateTime
				FROM 
					task as t,user as u,hostgroup as hg 
				WHERE
					t.createByID = u.id AND t.hostGroupID = hg.id and isDeleted=0`
	args := []interface{}{}
	var count int
	if len(ids) != 0 {
		getsql += " AND ("
		querys := []string{}
		for _, id := range ids {
			querys = append(querys, "t.id=?")
			args = append(args, id)
		}
		getsql += strings.Join(querys, " OR ")
		getsql += ")"
	}
	if name != "" {
		getsql += " AND t.name like '" + name + "%' "
	}
	if runStatus >= 0 {
		getsql += " AND t.run =?"
		args = append(args, runStatus)
	}
	if createbyid != "" {
		getsql += " AND t.createByID=?"
		args = append(args, createbyid)
	}
	getsql += " order by t.createTime desc"
	tasks := []define.GetTask{}
	if limit > 0 {
		var err error
		count, err = countColums(ctx, getsql, args...)
		if err != nil {
			return tasks, 0, fmt.Errorf("countColums failed: %w", err)
		}
		getsql += " LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}
	conn, err := db.GetConn(ctx)
	if err != nil {
		return tasks, 0, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()

	// slog.Println(slog.DEBUG, getsql)
	stmt, err := conn.PrepareContext(ctx, getsql)
	if err != nil {
		return tasks, 0, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return tasks, 0, fmt.Errorf("stmt.QueryContext failed: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		t := define.GetTask{}
		var (
			parentTaskIds, childTaskIds string
			alarmUserids                string
		)

		err = rows.Scan(&t.ID,
			&t.Name,
			&t.Group,
			&t.TaskType,
			&t.Ip,
			&t.Port,
			&t.ProbeId,
			&t.Plug,
			&t.Run,
			&t.Status,
			&parentTaskIds,
			&t.ParentRunParallel,
			&childTaskIds,
			&t.ChildRunParallel,
			&t.Cronexpr,
			&t.Priority,
			&t.Timeout,
			&t.Threads,
			&alarmUserids,
			&t.RoutePolicy,
			&t.AlarmStatus,
			&t.CreateBy,
			&t.CreateByUID,
			&t.HostGroup,
			&t.HostGroupID,
			&t.Remark,
			&t.CreateTime,
			&t.UpdateTime,
		)
		if err != nil {
			log.Error("rows.Scan ", zap.Error(err))
			continue
		}
		t.AlarmUserIds = []string{}
		if alarmUserids != "" {
			t.AlarmUserIds = append(t.AlarmUserIds, strings.Split(alarmUserids, ",")...)
			users, _, err := GetUsers(ctx, t.AlarmUserIds, 0, 0)
			if err != nil {
				log.Error("GetUsers ids failed", zap.Strings("uids", t.AlarmUserIds))
			}
			t.AlarmUserIdsDesc = make([]string, 0, len(t.AlarmUserIds))
			for _, user := range users {
				t.AlarmUserIdsDesc = append(t.AlarmUserIdsDesc, user.Name)
			}
		}
		t.ParentTaskIds = []string{}
		t.ParentTaskIdsDesc = []string{}
		if parentTaskIds != "" {
			t.ParentTaskIds = append(t.ParentTaskIds, strings.Split(parentTaskIds, ",")...)
			if first {
				ptasks, _, err := getTasks(ctx, t.ParentTaskIds, "", 0, 0, false, -1, "")
				if err != nil {
					log.Error("getTasks failed", zap.Error(err))
				}
				for _, task := range ptasks {
					t.ParentTaskIdsDesc = append(t.ParentTaskIdsDesc, task.Name)
				}
			}

		}
		t.ChildTaskIds = []string{}
		t.ChildTaskIdsDesc = []string{}
		if childTaskIds != "" {
			t.ChildTaskIds = append(t.ChildTaskIds, strings.Split(childTaskIds, ",")...)
			if first {
				ctasks, _, err := getTasks(ctx, t.ChildTaskIds, "", 0, 0, false, -1, "")
				if err != nil {
					log.Error("getTasks failed", zap.Error(err))
				}
				for _, task := range ctasks {
					t.ChildTaskIdsDesc = append(t.ChildTaskIdsDesc, task.Name)
				}
			}

		}
		if err != nil {
			log.Error("GetDataRun failed", zap.Any("type", t.TaskType), zap.Error(err))
			continue
		}
		t.RoutePolicyDesc = t.RoutePolicy.String()
		t.TaskTypeDesc = t.TaskType.String()
		t.AlarmStatusDesc = t.AlarmStatus.String()
		t.StatusDesc = t.Status.String()
		tasks = append(tasks, t)
	}
	return tasks, count, nil
}

func UpdateTaskStatus(ctx context.Context, taskId string, run int, status define.TaskOneStatus) error {
	updatesql := `UPDATE task set run=?,status=? WHERE id=?`
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
	result, err := stmt.ExecContext(ctx, run, status, taskId)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	line, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected failed: %w", err)
	}
	if line <= 0 {
		slog.Println(slog.DEBUG, "修改失败")
	}
	return nil
}
