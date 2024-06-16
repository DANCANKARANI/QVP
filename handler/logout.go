package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	//"main.go/middleware"
	"main.go/model"
	"main.go/utilities"
)
func Logout(c *fiber.Ctx) error {
	tokenString,err :=utilities.GetJWTToken(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusUnauthorized)
	}
	fmt.Println(tokenString)
	
	err =model.InvalidateToken(tokenString)
	if err != nil {
		return utilities.ShowError(c,"failed to invalidate the token",fiber.StatusUnauthorized)
	}
	return utilities.ShowMessage(c,"successfully logged out",fiber.StatusOK)
}