package models

import (
	"database/sql"
	"fmt"
	"simpleapp/database"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB = database.MakeConnection()

type Admin struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `'json:"password"`
}

func (a Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))

	return err == nil
}

func (Admin) All() ([]Admin, error) {
	var admins []Admin

	rows, err := db.Query("SELECT * FROM admins")
	if err != nil {
		return admins, fmt.Errorf("get all admins: %v", err)
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

func (a *Admin) Save() (int64, error) {
	passwd, err := bcrypt.GenerateFromPassword([]byte(a.Password), 14)
	if err != nil {
		return 0, fmt.Errorf("Save Admin: %v", err)
	}

	res, err := db.Exec("insert into admins (name, username, password) values(?, ?, ?)", a.Name, a.Username, passwd)
	if err != nil {
		return 0, fmt.Errorf("Save Admin: %v", err)
	}

	return res.LastInsertId()
}

func FindAdmin(id int64) (Admin, error) {
	var admin Admin

	row := db.QueryRow("select * from admins where id = ?", id)

	err := row.Scan(&admin.ID, &admin.Name, &admin.Username, &admin.Password)
	if err != nil {
		return admin, fmt.Errorf("find admin: %v", err)
	}

	return admin, nil
}

func FindAdminByUsername(username string) (Admin, error) {
	var admin Admin

	row := db.QueryRow("select * from admins where username = ?", username)

	err := row.Scan(&admin.ID, &admin.Name, &admin.Username, &admin.Password)
	if err != nil {
		return admin, err
	}

	return admin, nil
}
