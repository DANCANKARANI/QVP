package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/user_handler"
)

func EndPoints() {
	app := fiber.New()
	app.Post("/register",user_handler.CreateUserAccount)
	app.Post("/login",user_handler.Login)
	app.Post("/forgot-password",user_handler.ForgotPassword)
	app.Post("/reset-password",user_handler.ResetPassword)
	//start the server

	//JWT middleware
	app.Use(user_handler.JWTMiddleware)
	
	//protected routes
	app.Post("register-dependant", user_handler.RegisterDependantAccount)
	app.Get("get-dependants",user_handler.GetDependantsHandler)
	app.Get("/logout",user_handler.Logout)

	app.Listen(":3000")
}
