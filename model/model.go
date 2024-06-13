package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct{
	ID      uuid.UUID 				`gorm:"type:char(36);primary_key"`
	FullName 		string 			`json:"full_name" gorm:"size:255"`
	Email 		string 				`json:"email" gorm:"size:255"`
	PhoneNumber string 				`json:"phone_number" gorm:"size:255"`
	CountryCode string 				`json:"country_code" gorm:"size:10"`
	Password 	string 				`json:"password" gorm:"size:255"`
	ConfirmPassword string 			`json:"confirm_password" gorm:"size:255"`
	ResetCode	string				`json:"reset_code"`
	CodeExpirationTime time.Time 	`json:"code_expiration_time"`
}

