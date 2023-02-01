package AdminController

import (
	"encoding/json"
	"net/http"
	"simpleapp/models"
)

func List (w http.ResponseWriter, r *http.Request) {
	adminModel := models.Admin{}

	adminsList, _ := adminModel.All()

	err := json.NewEncoder(w).Encode(adminsList)
	if err != nil {
		panic(err.Error())
	}
}
