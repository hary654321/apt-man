package models

import (
	"fmt"
	"strings"

	"zrDispatch/core/config"
	"zrDispatch/core/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

func Setup() {
	var err error

	slog.Println(slog.DEBUG, config.CoreConf.Server.DB.Drivename, config.CoreConf.Server.DB.Dsn)
	db, err = gorm.Open(mysql.Open(config.CoreConf.Server.DB.Dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		slog.Printf(slog.ERROR, "models.Setup err: %v", err)
	}

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return config.CoreConf.Server.DB.TablePrefix + "_" + defaultTableName
	// }

}

func getInWhere(conds []string) string {

	ss := strings.Join(conds, "','")
	res := fmt.Sprintf("(%s)", ss)

	println(res)
	return res
}
