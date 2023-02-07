package AuthController

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"simpleapp/models"
	"simpleapp/modules/jwtHandler"
	"simpleapp/validations/login"
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

	if err := login.ValidateUsername(username); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("validation error: %v", err.Error()),
		})

		return
	}

	if err := login.ValidatePassword(password); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("validation error: %v", err.Error()),
		})

		return
	}

	admin, err := models.FindAdminByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "invalid username or password!",
		})

		return
	}

	if !admin.CheckPassword(password) {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "invalid username or password!",
		})

		return
	}

	jwt, err := jwtHandler.Generate(admin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Internal Server Error!",
		})
		
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": jwt,
	})

	// or we can set cookie
	/* http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwt,
		Expires: expireTime,
	}) */
}
