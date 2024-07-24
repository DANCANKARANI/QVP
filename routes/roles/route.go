package roles

import (
	"github.com/DANCANKARANI/QVP/controllers/role"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetRoleRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/role")
	roleGroup := auth.Group("/",user.JWTMiddleware)
	roleGroup.Post("/",role.AddRoleHandler)
	roleGroup.Get("/",role.GetRolesHandler)
	roleGroup.Patch("/:id",role.UpdateRoleHandler)
	roleGroup.Delete("/:id",role.DeleteRoleHandler)
	roleGroup.Post("/permission",role.AssociatePermissionsHandler)
}