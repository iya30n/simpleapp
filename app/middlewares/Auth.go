package middlewares

import (
	"net/http"
	"simpleapp/app/modules/jwtHandler"
	"simpleapp/core/responseHandler"
)

type Auth struct{}

func (Auth) Handle(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("Authorization")
	if len(token) < 1 || !jwtHandler.Verify(token) {
		responsehandler.Json(w, map[string]string{"message": "UnAuthenticated!"}, http.StatusBadRequest)
		return false
	}
	
	return true
}
