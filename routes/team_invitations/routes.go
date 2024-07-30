package team_invitations

import (
	"github.com/DANCANKARANI/QVP/controllers/team_invitation"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetTeamInvitationRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/team-invitations")
	invitationGroup := auth.Group("/", user.JWTMiddleware)
	invitationGroup.Post("/",team_invitation.AddTeamInvitationHandler)
	invitationGroup.Get("/",team_invitation.GetTeamInvitationsHandler)
	invitationGroup.Get("/:id",team_invitation.GetTeamInvitationByTeamHandler)
	invitationGroup.Patch("/:id",team_invitation.UpdateTeamInvitationHandler)
	invitationGroup.Delete("/:id",team_invitation.DeleteTeamInvitationHandler)
}