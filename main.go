package main

import (
	"fmt"
	"log"
	"simpleapp/models"
)

func main() {
	// admins, err := models.Admin{}.All()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// for _, admin := range admins {
	// 	fmt.Println(admin.Name)
	// }

	admin := &models.Admin{
		Name: "mamad",
		Username: "mamad",
		Password: "3421",
	}

	adminId, err := admin.Save()
	if err != nil {
		log.Fatal(err.Error())
	}

	savedAdmin, err := models.FindAdmin(adminId)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(savedAdmin.Name)
}
