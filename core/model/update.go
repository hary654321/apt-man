package model

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"zrDispatch/common/db"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/asset"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"go.uber.org/zap"
)

func Update() {

	fs := &assetfs.AssetFS{
		Asset:     asset.Asset,
		AssetDir:  asset.AssetDir,
		AssetInfo: asset.AssetInfo,
	}

	sqlfilename := "sql/update.sql"
	file, err := fs.Open(sqlfilename)
	if err != nil {
		slog.Println(slog.DEBUG, "fs.Open failed", zap.String("filename", sqlfilename), zap.Error(err))
		return
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		slog.Println(slog.DEBUG, "ioutil.ReadAll failed", zap.Error(err))
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

	for _, sql := range strings.Split(string(content), ";") {
		_, err = conn.ExecContext(context.Background(), sql)
		if err != nil {
			slog.Println(slog.DEBUG, "Update", "sql", sql, err)
			continue
		}
	}
	os.Remove("update.sql")

}
