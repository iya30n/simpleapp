package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Admin struct {
	ID       int64
	Name     string
	Username string
	Password string
}

func getAdmins() ([]Admin, error) {
	var admins []Admin

	rows, err := db.Query("SELECT * FROM admins")
	if err != nil {
		return admins, fmt.Errorf("getAdmins: %v", err)
	}

	for rows.Next() {
		var admin Admin
		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Username, &admin.Password); err != nil {
			return admins, fmt.Errorf("getAdmins: %v", err)
		}

		admins = append(admins, admin)
	}

	return admins, nil
}

var db *sql.DB

func main() {
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

	admins, err := getAdmins()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, admin := range admins {
		fmt.Println(admin.Name)
	}
}
