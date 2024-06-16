package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/handler"
)

func EndPoints() {
	app := fiber.New()
	app.Post("/register",handler.CreateUserAccount)
	app.Post("/login",handler.Login)
	app.Post("/forgot-password",handler.ForgotPassword)
	app.Post("/reset-password",handler.ResetPassword)
	//start the server

	//JWT middleware
	app.Use(handler.JWTMiddleware)
	
	//protected routes
	app.Get("/logout",handler.Logout)
	app.Listen(":3000")
}
