package users

import (
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/DANCANKARANI/QVP/database"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/gofiber/fiber/v2"
)
var db =database.ConnectDB()
func SetUserRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/user")
	auth.Post("/",user.CreateUserAccount)
	auth.Post("/login",user.Login)
	//protected routes
	userGroup := auth.Group("/",user.JWTMiddleware)
	userGroup.Use(model.SetUserIDMiddleware(db))
	userGroup.Get("/all",user.GetAllUsersHandler)
	userGroup.Get("/",user.GetOneUserHandler)
	userGroup.Post("/:id",user.AddUserInsurance)
	userGroup.Patch("/:id",user.UpdateUserInsurance)
	userGroup.Patch("/",user.UpdateUserHandler)
	userGroup.Post("/forgot-password",user.ForgotPassword)
	userGroup.Post("/reset-password",user.ResetPassword)
	userGroup.Get("/logout",user.Logout)
	userGroup.Patch("/:id",user.UpdateUserHandler)
}