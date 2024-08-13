package admins

import (
	"github.com/DANCANKARANI/QVP/controllers/admin"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetAdminsRoutes(app *fiber.App) {
	// Group routes under /api/v1/admin
	auth := app.Group("/api/v1/admins")
	auth.Post("/signup",admin.RegisterAdminHandler)
	auth.Post("/login",admin.AdminLogin)

	//protected routes
	adminGroup := auth.Group("/", user.JWTMiddleware)
	adminGroup.Get("/",admin.GetAdminDetailsHandler)
	adminGroup.Post("/logout",admin.Logout)
	adminGroup.Patch("/",admin.UpdateAdminHandler)
	adminGroup.Put("/",admin.UpdateProfilePicHandler)
	adminGroup.Get("/all",admin.GetAllAdminsHandler)
	adminGroup.Delete("/",admin.DeleteAdmin)

}