package Admin

import "fmt"

func Update(a AdminModel) (AdminModel, error) {
	passwd := a.Password

	_, err := DB.Exec(fmt.Sprintf("UPDATE admins SET name = %s, username = %s, password = %s", a.Name, a.Username, passwd))
	if err != nil {
		return a, fmt.Errorf("Update Admin: %v", err)
	}

	return a, nil
}
