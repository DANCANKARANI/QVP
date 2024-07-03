package model

import (
	"errors"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
gets user using user ID
*/
type ResponseUser struct{
	ID uuid.UUID 		`json:"id"`
	FullName string 	`json:"full_name"`
	PhoneNumber string 	`json:"phone_number"`
	Email string 		`json:"email"`
}

func GetOneUSer(c *fiber.Ctx)(*ResponseUser,error){
	id,err:=GetAuthUserID(c)
	if err != nil {
		return nil,errors.New("failed to get user's id:"+err.Error())
	}
	user := ResponseUser{}
	err = db.First(&User{},"id = ?",id).Scan(&user).Error
	if err != nil {
		return nil,errors.New("failed to get user details:"+err.Error())
	}
	return &user,nil
}
//gets all the users
func GetAllUsers(c *fiber.Ctx)(*[]ResponseUser,error){
	response:=[]ResponseUser{}
	err := db.Model(&User{}).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to get users:"+err.Error())
	}
	return &response,nil
}
//updates the user by id
func UpdateUser(c *fiber.Ctx)(*ResponseUser,error){
	id,err:=GetAuthUserID(c)
	if err != nil {
		return nil,errors.New("failed to get user's id:"+err.Error())
	}
	body := User{}
	imageURL,err := utilities.GenerateImageUrl(c)
	if err != nil {
		return nil, err
	}
	if err = c.BodyParser(&body);err != nil {
		return nil,errors.New("failed to parse:"+err.Error())
	}
	response := ResponseUser{}
	body.ProfilePhotoPath=imageURL
	err = db.First(&User{},"id = ?",id).Updates(&body).Scan(&response).Error
	if err != nil {
		return nil,errors.New("error in updating the user:"+err.Error())
	}
	return &response,nil
}

func AddProfileImage(c *fiber.Ctx)error{
	id,err:=GetAuthUserID(c)
	if err != nil {
		return errors.New("failed to get user's id:"+err.Error())
	}
	imageURL,err := utilities.GenerateImageUrl(c)
	if err != nil {
		return err
	}
	body := User{}
	body.ProfilePhotoPath=imageURL
	err = db.First(&User{},"id = ?",id).Updates(&body).Error
	if err != nil {
		return errors.New("failed to add profile")
	}
	return nil
}