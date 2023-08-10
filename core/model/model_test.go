package model

import (
	"testing"

	"zrDispatch/core/config"
	mylog "zrDispatch/core/utils/log"
)

func Test_countColums(t *testing.T) {
	querysql := `SELECT id,name,role,forbid,hashpassword FROM user`
	wantsql := `SELECT count() FROM user`
	gensql := gencountsql(querysql)
	if gensql != wantsql {
		t.Errorf("generate sql failed want getsql '%s',but gensql is '%s'", wantsql, gensql)
	}
}

func Test_ShowTable(t *testing.T) {
	config.Init("/Users/labulakalia/workerspace/golang/crocodile/core/config/core.toml")
	mylog.Init()
	InitDb()
	InitRabc()
	// conn,err := db.GetConn(context.Background())
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// rows,err := conn.QueryContext(context.Background(), "SELECT name FROM sqlite_master WHERE type ='table'")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// for rows.Next() {
	// 	var table string
	// 	rows.Scan(&table)
	// 	t.Log(table)
	// }

	enforcer := GetEnforcer()
	pass, err := enforcer.Enforce("238397974042906624", "/api/v1/hostgroup", "POST")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pass)
}
