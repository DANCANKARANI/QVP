package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/handler"
)

func EndPoints() {
	app := fiber.New()
	app.Post("/register",handler.CreateUserAccount)
	app.Post("login",handler.Login)

	//start the server
	app.Listen(":3000")
}
