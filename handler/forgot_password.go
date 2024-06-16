package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"main.go/model"
	"main.go/utilities"
)

type User struct{
	FullName string `json:"full_name"`
	Email string 		`json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func ForgotPassword(c *fiber.Ctx) error {
	
	user := User{}
	if err := c.BodyParser(&user);err!=nil {
		return utilities.ShowError(c,"failed to parse JSON data",fiber.StatusInternalServerError)
	}

	//checking if the user with the given email and phone number exists
	found_user,err :=model.FindUser(user.Email,user.PhoneNumber)
	if err != nil {
		return utilities.ShowError(c, "user does not exist",fiber.StatusNotFound)
	}

	//generate code and expiration time
	code,exp_time:=utilities.GenerateCode()
	err = model.AddResetCode(c,user.PhoneNumber,user.Email,code,exp_time)
	if err !=  nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}

	//send the code password reset code
	err =utilities.SendEmail(user.Email,code,exp_time)
	if err != nil {
		fmt.Println(err.Error())
	}
	return utilities.ShowSuccess(c,"link sent to your email",fiber.StatusOK,User{found_user.FullName,found_user.Email,found_user.PhoneNumber})
}

