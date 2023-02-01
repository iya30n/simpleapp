package AuthController

import (
	//"encoding/json"
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
	// TODO: validate username, password

	username, password := r.PostFormValue("username"), r.PostFormValue("password")
	
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
