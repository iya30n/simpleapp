package AdminController

import (
	"net/http"
	"simpleapp/app/models/Admin"
	responsehandler "simpleapp/core/responseHandler"
)

func List(w http.ResponseWriter, r *http.Request) {
	adminsList, _ := Admin.All()

	responsehandler.Json(w, adminsList, http.StatusOK)
}
