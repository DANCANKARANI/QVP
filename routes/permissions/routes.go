package permissions

import (
	"github.com/DANCANKARANI/QVP/controllers/permission"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetPermissionRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/permission")
	permissionGroup := auth.Group("/",user.JWTMiddleware)
	permissionGroup.Post("/",permission.AddPermissionHandler)
	permissionGroup.Get("/",permission.GetPermissionHandler)
	permissionGroup.Patch("/:id",permission.UpdatePermissionHandler)
	permissionGroup.Delete("/:id",permission.DeletePermissionHandler)
}