package insurance_users

import (
	"github.com/DANCANKARANI/QVP/controllers/insurance_user"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetInsuranceUserRoutes(app *fiber.App) {
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/insurance-users")
	insUserGroup := auth.Group("/", user.JWTMiddleware)
	insUserGroup.Post("/",insurance_user.AssignInsuranceUserHandler)
	insUserGroup.Patch("/:id",insurance_user.UpdateInsuranceUserHandler)
	insUserGroup.Delete("/:id",insurance_user.DeleteInsuranceUserHandler)
	insUserGroup.Get("/",insurance_user.GetUserWithInsuranceHandler)
}