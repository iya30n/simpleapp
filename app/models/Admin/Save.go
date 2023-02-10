package Admin

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Save(a AdminModel) (int64, error) {
	passwd, err := bcrypt.GenerateFromPassword([]byte(a.Password), 14)
	if err != nil {
		return 0, fmt.Errorf("Save Admin: %v", err)
	}

	res, err := DB.Exec("insert into admins (name, username, password) values(?, ?, ?)", a.Name, a.Username, passwd)
	if err != nil {
		return 0, fmt.Errorf("Save Admin: %v", err)
	}

	return res.LastInsertId()
}
