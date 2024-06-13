package model

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/utilities"
)
var db =database.ConnectDB()
//find user using phone number only
func UserExist(c *fiber.Ctx,phone_number string)(bool,error,User){
	 existingUser := User{}
	result := db.Where("phone_number = ?",phone_number).First(&existingUser)
	if result.Error != nil {
		//user not found
		return false,result.Error,existingUser
	}
	return true, nil,existingUser
}

func AddCode(c *fiber.Ctx,phone_number,email,code string,exp_time time.Time) error {
	user := User{}
	db.AutoMigrate(&user)
	result:=db.Where("phone_number = ? AND email = ?",phone_number,email).First(&user)
	if result.Error != nil {
		return utilities.ShowError(c,"failed to get user",fiber.StatusInternalServerError)
	}
	user.ResetCode=code
	user.CodeExpirationTime=exp_time
	result = db.Save(&user)
	if result.Error != nil {
		return utilities.ShowError(c,"failed to save data",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"code sent",fiber.StatusOK)
}

func FindUser(email, phone_number string)(User,error){
	user := User{}
	result:=db.Where("phone_number = ? AND email = ?",phone_number,email).First(&user)
	if result.Error != nil {
		return user,result.Error
	}
	return user,nil
}