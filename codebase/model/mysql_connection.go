package model

import (
	"fmt"
	"database/sql"

	"wc-accounting-go/codebase/cfg"

	"github.com/beego/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(configs *cfg.Configs, env string) {
	var connString string
	createTables(configs)
	if env == "dev" {
		connString = mysqlDriverURI(configs.Dev.Mysql.DBRootUser, configs.Dev.Mysql.DBHost, configs.Dev.Mysql.DBRootPassword, configs.Dev.Mysql.Port, configs.Dev.Mysql.DBName)
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connString)
}


func createTables(configs *cfg.Configs){
	var err error
	createdbStmt := "CREATE DATABASE IF NOT EXISTS wecare;"
	mysqlUri := mysqlDriverURI(configs.Dev.Mysql.DBRootUser, configs.Dev.Mysql.DBHost, configs.Dev.Mysql.DBRootPassword, configs.Dev.Mysql.Port, "")
	db, err := sql.Open("mysql", mysqlUri)
	defer db.Close()
	if _, err = db.Exec(createdbStmt); err != nil{
		fmt.Println(err)
		return
	}
	db.Close()
	mysqlUri = mysqlDriverURI(configs.Dev.Mysql.DBRootUser, configs.Dev.Mysql.DBHost, configs.Dev.Mysql.DBRootPassword, configs.Dev.Mysql.Port, configs.Dev.Mysql.DBName)
	db, err = sql.Open("mysql", mysqlUri)
	for _, query := range createTableQueries{
		_, err = db.Exec(query)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Construct a mysql protocol URI. The <dbname> parameter should be empty string when creating/dropping databases.
// If <dbname> is a valid database name, then the returned URI is configured to `USE <dbname>;` upon opening the db.
func mysqlDriverURI(user, host, password string, port int, dbname string) (uri string) {
	// check input.
	if user == "" || host == "" {
		panic("Params <user>, <host>, and <port> must ")
	} else if port < 1 || port > 65535 {
		panic(fmt.Sprintf("Invalid <port> number: %d", port))
	}

	var dbsuffix string
	if dbname != "" {
		// only set the suffix when the caller specifies a database.
		dbsuffix = fmt.Sprintf("%s?interpolateParams=true&parseTime=true", dbname)
	}
	uri = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbsuffix)
	return
}