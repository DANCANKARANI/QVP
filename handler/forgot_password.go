package handler

import (
	"github.com/gofiber/fiber/v2"
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
	result := db.Where("email =? AND phone_number = ?",user.Email,user.PhoneNumber).First(&user)
	if result.Error != nil {
		return utilities.ShowError(c, "user does not exist",fiber.StatusForbidden)
	}
	return utilities.ShowSuccess(c,"link sent to your email",fiber.StatusOK,user)
}