package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/DANCANKARANI/QVP/database"
	"github.com/DANCANKARANI/QVP/utilities"
)
var db =database.ConnectDB()
/*
finds user using phone number only
@params phone_number
*/
func UserExist(c *fiber.Ctx,phone_number string)(bool,User,error){
	db.AutoMigrate(User{})
	 existingUser := User{}
	err := db.Find(&User{}, "phone_number = ?",phone_number).Scan(&existingUser).Error
	if err != nil {
		//user not found
		fmt.Println(phone_number)
		return false,existingUser,errors.New("user not found:"+err.Error())
	}
	
	return true,existingUser, nil
}
/*
updates the reset password code in the database
@params phone_number
@params email
@params reset_code
@paarams expiration_time
*/
func AddResetCode(c *fiber.Ctx,phone_number,email,code string,exp_time time.Time) error {
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
/*
finds if the user with the given email and phone number is registered
@params email
@params phone_number
*/
func FindUser(email, phone_number string)(User,error){
	user := User{}
	result:=db.Where("phone_number = ? AND email = ?",phone_number,email).First(&user)
	if result.Error != nil {
		return user,result.Error
	}
	return user,nil
}

/*
finds dependants using phone number only
@params phone_number
*/
func GetAuthUserID(c *fiber.Ctx)(uuid.UUID,error){
	user_id :=c.Locals("user_id")
	id,ok := user_id.(*uuid.UUID)
	if !ok{
		return uuid.Nil,errors.New("failed converting user_id to uuid")
	}
	user_id=*id
	return user_id.(uuid.UUID),nil
}