package users

import (
	"github.com/DANCANKARANI/QVP/controllers/team_user"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/user")
	auth.Post("/",user.CreateUserAccount)
	auth.Post("/login",user.Login)
	//protected routes
	userGroup := auth.Group("/",user.JWTMiddleware)
	userGroup.Get("/all",user.GetAllUsersHandler)
	userGroup.Get("/",user.GetOneUserHandler)
	userGroup.Post("/:id",user.AddUserInsurance)
	userGroup.Patch("/:id",user.UpdateUserInsurance)
	userGroup.Patch("/",user.UpdateUserHandler)
	userGroup.Post("/forgot-password",user.ForgotPassword)
	userGroup.Post("/reset-password",user.ResetPassword)
	userGroup.Get("/logout",user.Logout)
	userGroup.Patch("/:id",user.UpdateUserHandler)

	//get teams for a user
	userGroup.Get("/:user_id/teams",team_user.GetTeamsForUserHandler)
}