package Admin

import "fmt"

func Update(a AdminModel) (AdminModel, error) {
	passwd := a.Password

	_, err := DB.Exec("UPDATE admins SET name = ?, username = ?, password = ? WHERE id = ?", a.Name, a.Username, passwd, a.ID)
	if err != nil {
		return a, fmt.Errorf("Update Admin: %v", err)
	}

	return a, nil
}
