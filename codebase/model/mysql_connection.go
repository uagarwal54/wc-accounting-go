package model

import (
	"database/sql"
	"fmt"

	"wc-accounting-go/codebase/cfg"

	"github.com/beego/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(env string) {
	var connStringWithDB, connStringWithoutDB string
	if env == "dev" {
		// For each env type, these both types of connString needs to be made
		connStringWithoutDB = mysqlDriverURI(cfg.ConfigInst.Dev.Mysql.DBRootUser, cfg.ConfigInst.Dev.Mysql.DBHost, cfg.ConfigInst.Dev.Mysql.DBRootPassword, cfg.ConfigInst.Dev.Mysql.Port, "")
		connStringWithDB = mysqlDriverURI(cfg.ConfigInst.Dev.Mysql.DBRootUser, cfg.ConfigInst.Dev.Mysql.DBHost, cfg.ConfigInst.Dev.Mysql.DBRootPassword, cfg.ConfigInst.Dev.Mysql.Port, cfg.ConfigInst.Dev.Mysql.DBName)
	}
	createTables(&cfg.ConfigInst, connStringWithoutDB, connStringWithDB)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connStringWithDB)
}

func createTables(configs *cfg.Configs, connStringWithoutDB, connStringWithDB string) {
	var err error
	createdbStmt := "CREATE DATABASE IF NOT EXISTS wecare;"
	db, err := sql.Open("mysql", connStringWithoutDB)
	defer db.Close()
	if _, err = db.Exec(createdbStmt); err != nil {
		fmt.Println(err)
		return
	}
	db, err = sql.Open("mysql", connStringWithDB)
	fmt.Print("Creating Tables.")
	for _, query := range createTableQueries {
		fmt.Print(".")
		_, err = db.Exec(query)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("")
	fmt.Print("Inserting configs into the config table.")
	insertDataIntoConfigTable(db)
	return
}

func insertDataIntoConfigTable(db *sql.DB) {
	insertQuery := "insert into `config` (configKey, configValue) values (\"%s\",\"%s\");"
	for cKey, cValue := range configMap {
		fmt.Print(".")
		_, err := db.Exec(fmt.Sprintf(insertQuery, cKey, cValue))
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
