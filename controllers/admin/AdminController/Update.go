package AdminController

import (
	"net/http"
	"simpleapp/models/Admin"
	responsehandler "simpleapp/modules/responseHandler"
	"simpleapp/validations/adminValidation"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

var response map[string]string

func Update(w http.ResponseWriter, r *http.Request) {
	adminId := strings.Replace(r.URL.Path, "/admin/admins/", "", 1)

	admin, err := Admin.Find(adminId)
	if err != nil {
		response = map[string]string{"message": "Record Not Found!"}
		responsehandler.Json(w, response, http.StatusNotFound)
		return
	}

	p := bluemonday.UGCPolicy()

	name, username, password := p.Sanitize(r.PostFormValue("name")), p.Sanitize(r.PostFormValue("username")), p.Sanitize(r.PostFormValue("password"))

	if err := validator(name, username, password); err != nil {
		response = map[string]string{"message": err.Error()}

		responsehandler.Json(w, response, http.StatusBadRequest)

		return
	}

	if len(name) > 0 {
		admin.Name = name
	}

	if len(username) > 0 {
		admin.Username = username
	}

	if len(password) > 0 {
		admin.Password = password
	}

	Admin.Update(admin)

	response = map[string]string{"message": "admin updated"}
	responsehandler.Json(w, response, http.StatusAccepted)
}

func validator(name string, username string, password string) error {
	if err := adminValidation.ValidateName(name); err != nil {
		return err
	}

	if err := adminValidation.ValidateUsername(username); err != nil {
		return err
	}

	if err := adminValidation.ValidatePassword(password); err != nil {
		return err
	}

	return nil
}
