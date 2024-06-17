package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/dependant_handler"
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
	app.Post("/register-dependant", dependant_handler.RegisterDependantAccount)
	app.Get("/get-dependants",dependant_handler.GetDependantsHandler)
	app.Get("get-dependant-id",dependant_handler.GetDependantID)
	app.Post("/update/dependants",dependant_handler.UpdateDependant)
	app.Get("/logout",user_handler.Logout)

	app.Listen(":3000")
}
