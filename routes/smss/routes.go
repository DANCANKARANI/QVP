package smss

import (
	"github.com/DANCANKARANI/QVP/controllers/sms"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetSmsRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/sms")
	smsGroup := auth.Group("/", user.JWTMiddleware)
	smsGroup.Post("/",sms.AddSmsHandler)
	smsGroup.Patch("/:id",sms.UpdateSmsHandler)
	smsGroup.Delete("/:id",sms.DeleteSmsHandler)
	smsGroup.Get("/", sms.GetUserSmsHandler)
}