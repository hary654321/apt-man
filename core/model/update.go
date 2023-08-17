package model

import (
	"context"
	"os"
	"zrDispatch/common/db"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"
)

func Update() {

	if !utils.PathExists("update.sql") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	conn, err := db.GetConn(ctx)
	if err != nil {
		slog.Println(slog.DEBUG, "db.GetConn failed: %w", err)
	}

	defer conn.Close()

	execsql, _ := utils.ReadLineData("update.sql")

	for _, sql := range execsql {
		_, err = conn.ExecContext(context.Background(), sql)
		if err != nil {
			slog.Println(slog.DEBUG, "Update", "sql", sql, err)
			continue
		}
	}
	os.Remove("update.sql")

}
