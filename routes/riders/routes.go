package riders

import (
	"github.com/DANCANKARANI/QVP/controllers/rider"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetRiderRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/rider")
	auth.Post("/signup",rider.CreateRiderAccount)
	auth.Post("/login",rider.RiderLogin)

	riderGroup :=auth.Group("/",user.JWTMiddleware)
	riderGroup.Get("/",rider.GetRiderHandler)
	riderGroup.Patch("/:id",rider.UpdateRiderHandler)
	riderGroup.Delete("/:id", rider.DeleteRiderHandler)
	riderGroup.Post("/logout",rider.Logout)
}