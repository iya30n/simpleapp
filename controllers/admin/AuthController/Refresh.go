package AuthController

import (
	"net/http"
	"simpleapp/modules/jwtHandler"
	responsehandler "simpleapp/modules/responseHandler"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	newToken, err := jwtHandler.Refresh(token)
	if err != nil {
		responsehandler.Json(w, map[string]string{"message": err.Error(),}, http.StatusBadRequest)
		return
	}

	responsehandler.Json(w, map[string]string{"token": newToken,}, http.StatusAccepted)
}