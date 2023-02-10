package Admin

import (
	"fmt"
)

func Find(id string) (AdminModel, error) {
	var admin AdminModel

	row := DB.QueryRow("select id, name, username from admins where id = ?", id)

	err := row.Scan(&admin.ID, &admin.Name, &admin.Username)
	if err != nil {
		return admin, fmt.Errorf("find admin: %v", err)
	}

	return admin, nil
}

func FindByUsername(username string) (AdminModel, error) {
	var admin AdminModel

	row := DB.QueryRow("select * from admins where username = ?", username)

	err := row.Scan(&admin.ID, &admin.Name, &admin.Username, &admin.Password)
	if err != nil {
		return admin, err
	}

	return admin, nil
}
