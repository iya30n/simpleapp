package AuthController

import (
	"net/http"
	"simpleapp/app/models/Admin"
	responsehandler "simpleapp/app/modules/responseHandler"
	"simpleapp/app/validations/adminValidation"

	"github.com/microcosm-cc/bluemonday"
)

var response map[string]string

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: check if admin is logged in and has access to this action.

	p := bluemonday.UGCPolicy()

	name, username, password := p.Sanitize(r.PostFormValue("name")), p.Sanitize(r.PostFormValue("username")), p.Sanitize(r.PostFormValue("password"))

	if err := adminValidation.ValidateName(name); err != nil {

		response = map[string]string{"Message": err.Error()}

		responsehandler.Json(w, response, http.StatusBadRequest)

		return
	}

	if err := adminValidation.ValidateUsername(username); err != nil {

		response = map[string]string{"Message": err.Error()}

		responsehandler.Json(w, response, http.StatusBadRequest)

		return
	}

	if err := adminValidation.ValidatePassword(password); err != nil {

		response = map[string]string{"Message": err.Error()}

		responsehandler.Json(w, response, http.StatusBadRequest)

		return
	}

	if checkUsernameExists(username) {
		response = map[string]string{"message": "Username already taken!"}

		responsehandler.Json(w, response, http.StatusBadRequest)

		return
	}

	admin := Admin.AdminModel{
		Name:     name,
		Username: username,
		Password: password,
	}

	_, err := Admin.Save(admin)
	if err != nil {
		response = map[string]string{"message": "Server Error!. please try later."}

		responsehandler.Json(w, response, http.StatusInternalServerError)

		return
	}

	response = map[string]string{"message": "Admin created successfully."}

	responsehandler.Json(w, response, http.StatusCreated)
}

func checkUsernameExists(username string) bool {
	admin, _ := Admin.FindByUsername(username)

	return admin.ID > 0
}
