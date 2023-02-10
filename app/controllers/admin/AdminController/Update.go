package AdminController

import (
	"net/http"
	"simpleapp/app/models/Admin"
	responsehandler "simpleapp/app/modules/responseHandler"
	"simpleapp/app/validations/adminValidation"

	"github.com/microcosm-cc/bluemonday"
)

var response map[string]interface{}

func Update(w http.ResponseWriter, r *http.Request) {
	adminId := r.URL.Path[len("/admins/admin/edit/"):]

	admin, err := Admin.Find(adminId)
	if err != nil {
		response = map[string]interface{}{"message": "Record Not Found!"}
		responsehandler.Json(w, response, http.StatusNotFound)
		return
	}

	p := bluemonday.UGCPolicy()

	name, username, password := p.Sanitize(r.PostFormValue("name")), p.Sanitize(r.PostFormValue("username")), p.Sanitize(r.PostFormValue("password"))

	if err := validator(name, username, password); err != nil {
		response = map[string]interface{}{"message": err.Error()}

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

	updatedAdmin, err := Admin.Update(admin)
	if err != nil {
		response = map[string]interface{}{"message": "Server Error!"}
		responsehandler.Json(w, response, http.StatusInternalServerError)
		return
	}

	response = map[string]interface{}{"message": "admin updated", "admin": updatedAdmin}
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
