package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/middleware"
	"github.com/DANCANKARANI/QVP/utilities"
)

func JWTMiddleware(c *fiber.Ctx) error {
// Check for token in cookies first
tokenString := c.Cookies("Authorization")

	if tokenString == ""{
		
		return utilities.ShowError(c,"missing or malformed JWT",fiber.StatusUnauthorized)
	}
	//validate the token
	claims,err :=middleware.ValidateToken(tokenString)
	if err != nil{
		return errors.New("Unauthorized:"+err.Error())
	}
	//store the userID 
	c.Locals("user_id",claims.UserID)
	return c.Next()
}
