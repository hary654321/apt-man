package model

import (
	"context"
	"io/ioutil"
	"strings"
	"zrDispatch/common/db"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/asset"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"go.uber.org/zap"
)

func Update() {

	filename := "update1130.sql"

	if utils.PathExists(filename) {

		slog.Println(slog.DEBUG, "update.sql is exists")
		return
	}

	sqlfilename := "sql/" + filename
	fs := &assetfs.AssetFS{
		Asset:     asset.Asset,
		AssetDir:  asset.AssetDir,
		AssetInfo: asset.AssetInfo,
	}

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

	utils.Write(filename, "")

}

func Pi() {

	filename := "pi.sql"

	// if utils.PathExists(filename) {

	// 	slog.Println(slog.DEBUG, "update.sql is exists")
	// 	return
	// }

	sqlfilename := "sql/" + filename
	fs := &assetfs.AssetFS{
		Asset:     asset.Asset,
		AssetDir:  asset.AssetDir,
		AssetInfo: asset.AssetInfo,
	}

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

	for _, sql := range strings.Split(string(content), ";end") {
		_, err = conn.ExecContext(context.Background(), sql)
		if err != nil {
			slog.Println(slog.DEBUG, "Update", "sql", sql, err)
			continue
		}
	}

	utils.Write(filename, "")

}
