package riders

import (
	"github.com/DANCANKARANI/QVP/controllers/rider"
	//"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetRiderRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/rider")
	auth.Post("/",rider.RegisterRider)
	auth.Get("/",rider.GetRiderHandler)
}