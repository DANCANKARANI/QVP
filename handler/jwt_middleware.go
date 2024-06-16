package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"main.go/middleware"
	"main.go/utilities"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get(fiber.HeaderAuthorization)
	if authHeader == "" {
		
		return utilities.ShowError(c,"missing or malformed JWT",fiber.StatusUnauthorized)
	}
	//split the token into "Bearer" and token
	tokenString :=strings.TrimSpace(strings.TrimPrefix(authHeader,"Bearer"))

	if tokenString == ""{
		
		return utilities.ShowError(c,"missing or malformed JWT",fiber.StatusUnauthorized)
	}
	fmt.Println("1 ",authHeader)
	fmt.Println("2 ",tokenString)
	//validate the token
	claims,err :=middleware.ValidateToken(tokenString)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusUnauthorized)
	}
	//store the userID 
	c.Locals("user_id",claims.UserID.String())
	return c.Next()
}