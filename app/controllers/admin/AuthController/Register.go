package AuthController

import (
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"simpleapp/app/models/Admin"
	responsehandler "simpleapp/app/modules/responseHandler"
	errorHelper "simpleapp/core/helpers/error"
	"simpleapp/core/validator"
)

var response map[string]string

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: check if admin is logged in and has access to this action.

	p := bluemonday.UGCPolicy()

	name, username, password := p.Sanitize(r.PostFormValue("name")), p.Sanitize(r.PostFormValue("username")), p.Sanitize(r.PostFormValue("password"))

	validationRules := map[string]string{
		"name":     "required|string|min:3|max:50",
		"username": "required|string|min:3|max:50",
		"password": "required|string|min:8|max:100",
	}

	if errors := validator.Validate(r, validationRules); errors != nil {
		responsehandler.Json(w,
			map[string][]string{"errors": errorHelper.Stringify(errors)},
			http.StatusBadRequest)

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
