package AuthController

import (
	"net/http"
	"simpleapp/models"
	responsehandler "simpleapp/modules/responseHandler"
	"simpleapp/validations/adminValidation"
)

var response map[string]string

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: prevent xss
	// TODO: check if admin is logged in and has access to this action.

	name, username, password := r.PostFormValue("name"), r.PostFormValue("username"), r.PostFormValue("password")

	if err := adminValidation.ValidateName(name); err != nil {

		response = map[string]string{"Message": err.Error(),}

		responsehandler.Json(w, response, http.StatusBadRequest)

		return
	}

	if err := adminValidation.ValidateUsername(username); err != nil {

		response = map[string]string{"Message": err.Error(),}

		responsehandler.Json(w, response, http.StatusBadRequest)
		
		return
	}

	if err := adminValidation.ValidatePassword(password); err != nil {

		response = map[string]string{"Message": err.Error(),}

		responsehandler.Json(w, response, http.StatusBadRequest)
		
		return
	}

	if checkUsernameExists(username) {
		response = map[string]string{"message": "Username already taken!"}

		responsehandler.Json(w, response, http.StatusBadRequest)

		return
	}

	admin := models.Admin{
		Name: name,
		Username: username,
		Password: password,
	}

	_, err := admin.Save()
	if err != nil {
		response = map[string]string{"message": "Server Error!. please try later."}

		responsehandler.Json(w, response, http.StatusInternalServerError)

		return
	}

	response = map[string]string{"message": "Admin created successfully."}

	responsehandler.Json(w, response, http.StatusCreated)
}

func checkUsernameExists(username string) bool {
	admin, _ := models.FindAdminByUsername(username)

	return admin.ID > 0
}