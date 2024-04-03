package main

import (
	"jwt-authentication/initializers"
	"jwt-authentication/models"
)

func main() {
	db, err := initializers.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
