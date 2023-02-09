package AdminController

import (
	"encoding/json"
	"net/http"
	"simpleapp/models/Admin"
)

func List (w http.ResponseWriter, r *http.Request) {
	adminsList, _ := Admin.All()

	err := json.NewEncoder(w).Encode(adminsList)
	if err != nil {
		panic(err.Error())
	}
}
