package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MakeConnection() *sql.DB {
	if db != nil {
		return db
	}

	conf := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
		DBName:               "simpleapp",
	}

	var err error
	db, err = sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	return db
}
