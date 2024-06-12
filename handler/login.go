package handler

import (
	"github.com/gofiber/fiber/v2"
	"main.go/model"
	"main.go/utilities"
	"main.go/middleware"
)
type response_user struct{
	FullName string 	`json:"full_name"`
	PhoneNumber string 	`json:"phone_number"`
	Email string 		`json:"email"`
}

func Login(c *fiber.Ctx)error{
	
	user := model.User{}
	if err := c.BodyParser(&user); err !=nil {
		return c.JSON(fiber.Map{"error":err.Error()})
	}
	//check if the user exist
	var existingUser model.User
	result := db.Where("phone_number = ?",user.PhoneNumber).First(&existingUser)
	if result.Error != nil {
		//user not found
		return utilities.ShowError(c,"user does not exist",fiber.StatusInternalServerError)
	}
		
	//compare password
	err :=utilities.CompareHashAndPassowrd(existingUser.Password,user.Password)
	if err !=nil{
		return utilities.ShowError(c,"invalid credintials",fiber.StatusForbidden)
			 
	}
	//generating token
	tokenString,err := middleware.GenerateJWT(c,existingUser.ID.String())
	if err != nil{
		return utilities.ShowError(c,"failed to generate token",fiber.StatusInternalServerError)
	}
	response_user:=response_user{
		FullName: existingUser.FullName,
		PhoneNumber: existingUser.PhoneNumber,
		Email: existingUser.Email,
	}

	//set the authorization header with the token
	c.Set("Authorization","Bearer"+tokenString)
	return utilities.ShowSuccess(c,"successfully logged in",fiber.StatusOK,response_user)	
}

