package modules

import (
	"github.com/DANCANKARANI/QVP/controllers/module"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetModulesRoutes(app *fiber.App) {
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/modules")
	moduleGroup := auth.Group("/", user.JWTMiddleware)
	moduleGroup.Post("/",module.CreateModuleHandler)
	moduleGroup.Get("/", module.GetModulesHandler)
	moduleGroup.Patch("/:id",module.UpdateModuleHandler)
	moduleGroup.Delete("/:id",module.DeleteModuleHandler)
}