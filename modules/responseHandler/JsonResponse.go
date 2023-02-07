package responsehandler

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)

	json.NewEncoder(w).Encode(data)
}