package AuthController

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"simpleapp/models"
)

// this solution works for raw json
/* type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var l loginData
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "your username is %s, your password is %s", l.Username, l.Password)
} */

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: prevent xss and sql injection

	username, password := r.PostFormValue("username"), r.PostFormValue("password")

	if err := validateUsername(username); err != nil {
		fmt.Fprintf(w, "validation error: %v", err.Error())
		return
	}

	if err := validatePassword(password); err != nil {
		fmt.Fprintf(w, "validation error: %v", err.Error())
		return
	}

	admin, err := models.FindAdminByUsername(username)
	if err != nil {
		fmt.Fprintf(w, "error on login: %v", err.Error())
		return
	}

	if !admin.CheckPassword(password) {
		fmt.Fprintln(w, "invalid username or password")
		return
	}

	fmt.Fprintln(w, "You are logged in!")
}

func validateUsername(username string) error {
	if len(username) < 3 {
		return fmt.Errorf("username should be more than 3 characters")
	}

	if len(username) > 100 {
		return fmt.Errorf("username should be less than 100 characters")
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password should be more than 8 characters")
	}

	if len(password) > 100 {
		return fmt.Errorf("password should be less than 100 characters")
	}

	return nil
}