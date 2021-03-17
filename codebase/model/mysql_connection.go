package model

import (
	"fmt"

	"wc-accounting-go/codebase/cfg"

	"github.com/beego/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(configs *cfg.Configs, env string) {
	var connString string
	if env == "dev" {
		connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			configs.Dev.Mysql.DBRootUser,
			configs.Dev.Mysql.DBRootPassword,
			configs.Dev.Mysql.DBHost,
			configs.Dev.Mysql.Port,
			configs.Dev.Mysql.DBName,
		)
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connString)

}
