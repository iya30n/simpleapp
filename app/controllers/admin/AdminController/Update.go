package AdminController

import (
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"simpleapp/app/models/Admin"
	responsehandler "simpleapp/app/modules/responseHandler"
	errorHelper "simpleapp/core/helpers/error"
	"simpleapp/core/validator"
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

	validationRules := map[string]string{
		"name":     "string|min:3|max:50",
		"username": "string|min:3|max:50",
		"password": "string|min:8|max:100",
	}

	if errors := validator.Validate(r, validationRules); errors != nil {
		responsehandler.Json(w,
			map[string][]string{"errors": errorHelper.Stringify(errors)},
			http.StatusBadRequest)

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
