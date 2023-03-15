package AuthController

import (
	"net/http"
	"simpleapp/app/modules/jwtHandler"
	"simpleapp/core/responseHandler"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	// Cookie example

	/* c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token := c.Value */

	token := r.Header.Get("Authorization")

	newToken, err := jwtHandler.Refresh(token)
	if err != nil {
		responsehandler.Json(w, map[string]string{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	responsehandler.Json(w, map[string]string{"token": newToken}, http.StatusAccepted)
}
