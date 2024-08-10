package teams

import (
	"github.com/DANCANKARANI/QVP/controllers/team"
	"github.com/DANCANKARANI/QVP/controllers/team_user"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetTeamRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/teams")
	teamGroup := auth.Group("/", user.JWTMiddleware)
	teamGroup.Post("/",team.CreateTeamHandler)
	teamGroup.Patch("/:id",team.UpdateTeamHandler)
	teamGroup.Get("/",team.GetTeamsHandler)
	teamGroup.Delete("/:id",team.DeleteTeamHandler)

	//team users
	teamGroup.Post("/:team_id/users/:user_id",team_user.AddUserToTeamHandler)
	teamGroup.Get("/:team_id/users",team_user.GetUsersFromTeamHandler)
	teamGroup.Delete("/:team_id/users/:user_id",team_user.RemoveUserFromTeamHandler)
	teamGroup.Patch("/:team_id/users/:user_id",team_user.UpdateUserInTeamHandler)
}