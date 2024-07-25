package team_invitations

import (
	"github.com/DANCANKARANI/QVP/controllers/team_invitation"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetTeamInvitatatRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/team-invitations")
	invitationGroup := auth.Group("/", user.JWTMiddleware)
	invitationGroup.Post("/",team_invitation.AddTeamInvitationHandler)
}