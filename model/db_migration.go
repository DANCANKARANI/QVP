package model

import "fmt"

//migrates the models
func Migration() {
	fmt.Println("me 1")
	db.AutoMigrate(
		&Image{},
		&User{}, 
		&Dependant{},
	 	&Insurance{},
		&PaymentMethod{},
		&Payment{},
		&Notification{},
		&Prescription{},
		&Admin{},
		&Rider{},
	)
}