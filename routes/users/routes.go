package users

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/user"
)

func SetUserRoutes(app *fiber.App) {
	
	auth := app.Group("/api/v1/user")
	auth.Post("/",user.CreateUserAccount)
	auth.Post("/login",user.Login)
	//protected routes
	userGroup := auth.Group("/",user.JWTMiddleware)
	userGroup.Get("/all",user.GetAllUsersHandler)
	userGroup.Get("/",user.GetOneUserHandler)
	userGroup.Post("/reset-password",user.ResetPassword)
	userGroup.Patch("/profile",user.AddProfileImage)
	userGroup.Get("/logout",user.Logout)
	userGroup.Patch("/:id",user.UpdateUserHandler)
}