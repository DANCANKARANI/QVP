package model

import "log"

//migrates the models
func Migration() {
	log.Println("starting migrations...")
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
		&InsuranceUser{},
		&Role{},
		&Permission{},
		&Module{},
		&Team{},
		&TeamInvitation{},
		&QuoteDetail{},
		&Audit{},
		&TeamUser{},
		&PrescriptionDetail{},
	)
}