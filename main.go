package main

import (
	"fmt"
	"log"
	"simpleapp/models"
)

func main() {
	admins, err := models.Admin{}.All()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, admin := range admins {
		fmt.Println(admin.Name)
	}
}
