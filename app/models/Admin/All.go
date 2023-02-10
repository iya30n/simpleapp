package Admin

import (
	"fmt"
)

func All() ([]AdminModel, error) {
	var admins []AdminModel

	rows, err := DB.Query("SELECT id, name, username FROM admins")
	if err != nil {
		return admins, fmt.Errorf("get all admins: %v", err)
	}

	for rows.Next() {
		var admin AdminModel
		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Username); err != nil {
			return admins, fmt.Errorf("getAdmins: %v", err)
		}

		admins = append(admins, admin)
	}

	return admins, nil
}
