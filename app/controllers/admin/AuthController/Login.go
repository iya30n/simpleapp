package AuthController

import (
	"net/http"
	"simpleapp/app/models/Admin"
	"simpleapp/app/modules/jwtHandler"
	errorHelper "simpleapp/core/helpers/error"
	"simpleapp/core/responseHandler"
	"simpleapp/core/validator"
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

var responseData map[string]string

func Login(w http.ResponseWriter, r *http.Request) {
	username, password := r.PostFormValue("username"), r.PostFormValue("password")

	validationRules := map[string]string{
		"username": "required|string|min:3|max:50",
		"password": "required|string|min:8|max:50",
	}

	if err := validator.Validate(r, validationRules); err != nil {
		responsehandler.Json(w,
			map[string][]string{"errors": errorHelper.Stringify(err)},
			http.StatusBadRequest)

		return
	}

	admin, err := Admin.FindByUsername(username)
	if err != nil {
		responseData = map[string]string{
			"message": "invalid username or password!",
		}

		responsehandler.Json(w, responseData, http.StatusBadRequest)

		return
	}

	if !admin.CheckPassword(password) {
		responseData = map[string]string{
			"message": "invalid username or password!",
		}

		responsehandler.Json(w, responseData, http.StatusBadRequest)

		return
	}

	jwt, err := jwtHandler.Generate(admin)
	if err != nil {
		responseData = map[string]string{
			"message": "Internal Server Error!",
		}

		responsehandler.Json(w, responseData, http.StatusInternalServerError)

		return
	}

	responseData = map[string]string{
		"token": jwt,
	}

	responsehandler.Json(w, responseData, http.StatusAccepted)

	// or we can set cookie
	/* http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwt,
		Expires: expireTime,
	}) */
}
