package model

import (
	"main.go/database"
)

//migrates the models
func Migration() {
	db := database.ConnectDB()
	db.AutoMigrate(
		&User{}, 
		&Dependant{},
	 	&Insurance{},
		&PaymentMethod{},
		&Payment{},
		&Notification{},
	)
}