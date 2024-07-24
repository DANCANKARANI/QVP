package teams

import (
	"github.com/DANCANKARANI/QVP/controllers/team"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetTeamRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/teams")
	teamGroup := auth.Group("/", user.JWTMiddleware)
	teamGroup.Post("/",team.CreateTeamHandler)
	teamGroup.Patch("/:id",team.UpdateTeamHandler)
	teamGroup.Get("/",team.GetTeamsHandler)
	teamGroup.Delete("/",team.DeleteTeamHandler)
}