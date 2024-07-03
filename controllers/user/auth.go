package user

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/middleware"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
)
type ResponseUser struct{
	FullName string 	`json:"full_name"`
	PhoneNumber string 	`json:"phone_number"`
	Email string 		`json:"email"`
}
type loginResponse struct {
	Token string `json:"token"`
}

func Login(c *fiber.Ctx)error{
	user := model.User{}
	if err := c.BodyParser(&user); err !=nil {
		return c.JSON(fiber.Map{"error":err.Error()})
	}

	//check of user exist
	userExist,existingUser,err:= model.UserExist(c,user.PhoneNumber)
	fmt.Println(user.PhoneNumber+"1",existingUser.PhoneNumber+"2")
	if ! userExist {
		return utilities.ShowError(c,err.Error(),fiber.StatusNotFound)
	}
	
	//compare password
	err =utilities.CompareHashAndPassowrd(existingUser.Password,user.Password)
	if err !=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusForbidden)
			 
	}
	exp :=time.Hour*24
	//generating token
	tokenString,err := middleware.GenerateToken(middleware.Claims{UserID: &existingUser.ID},exp)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	//set token cookie 
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24), // Same duration as the token
		HTTPOnly: true, // Important for security, prevents JavaScript access
		Secure:   true, // Use secure cookies in production
		Path:     "/",  // Make the cookie available on all routes
	})
	response_user:=loginResponse{
		Token: tokenString,
	}
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

