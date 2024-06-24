package user_handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"main.go/middleware"
	"main.go/model"
	"main.go/utilities"
)
type ResponseUser struct{
	FullName string 	`json:"full_name"`
	PhoneNumber string 	`json:"phone_number"`
	Email string 		`json:"email"`
}

func Login(c *fiber.Ctx)error{
	user := model.User{}
	if err := c.BodyParser(&user); err !=nil {
		return c.JSON(fiber.Map{"error":err.Error()})
	}

	//check of user exist
	userExist,existingUser,_:= model.UserExist(c,user.PhoneNumber)
	if ! userExist {
		return utilities.ShowError(c,"user does not exist",fiber.StatusNotFound)
	}
	
	//compare password
	err :=utilities.CompareHashAndPassowrd(existingUser.Password,user.Password)
	if err !=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusForbidden)
			 
	}
	exp :=time.Hour*24
	//generating token
	//tokenString,err := middleware.GenerateJWT(c,existingUser.ID.String())
	tokenString,err := middleware.GenerateToken(middleware.Claims{UserID: &existingUser.ID},exp)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	response_user:=ResponseUser{
		FullName: existingUser.FullName,
		PhoneNumber: existingUser.PhoneNumber,
		Email: existingUser.Email,
	}

	//set the authorization header with the token
	c.Set("Authorization",tokenString)
	return utilities.ShowSuccess(c,"successfully logged in",fiber.StatusOK,response_user)	
}


func Logout(c *fiber.Ctx) error {
	tokenString,err :=utilities.GetJWTToken(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusUnauthorized)
	}
	fmt.Println(tokenString)
	err = middleware.InvalidateToken(tokenString)
	if err != nil {
		return utilities.ShowError(c,"failed to invalidate the token",fiber.StatusInternalServerError)
	}
	
	return utilities.ShowMessage(c,"successfully logged out",fiber.StatusOK)
}

