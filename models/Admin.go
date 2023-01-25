package models

import (
	"fmt"
	"simpleapp/database"
)

type Admin struct {
	ID       int64
	Name     string
	Username string
	Password string
}

func (Admin) All() ([]Admin, error) {
	db := database.MakeConnection()

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
