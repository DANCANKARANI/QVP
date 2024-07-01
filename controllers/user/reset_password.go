package user

import (
	//"time"

	"github.com/gofiber/fiber/v2"
	"main.go/model"
	"main.go/password"
	"main.go/utilities"
)

func ResetPassword(c *fiber.Ctx)error{
	user := model.User{}
	if err := c.BodyParser(&user); err != nil{
		return utilities.ShowError(c,"failed to parse Json data",fiber.StatusInternalServerError)
	}

	//call reset password
	password.ResetPassword(c,user.Email,user.PhoneNumber,user.Password,user.ResetCode)
	return utilities.ShowMessage(c,"password is changed succefully",fiber.StatusOK)
}