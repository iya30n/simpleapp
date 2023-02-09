package Admin

import (
	"golang.org/x/crypto/bcrypt"
)

type AdminModel struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `'json:"password"`
}

func (a AdminModel) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))

	return err == nil
}

