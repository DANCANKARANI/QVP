package dependants

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/controllers/dependant"
	"github.com/DANCANKARANI/QVP/controllers/user"
)

func SetDependantRoutes(app *fiber.App){
	// Group routes under /api/v1/dependants
	auth:=app.Group("/api/v1/dependants")
	dependantGroup := auth.Group("/",user.JWTMiddleware)
	dependantGroup.Get("/",dependant.GetDependantsHandler)
	dependantGroup.Post("/",dependant.RegisterDependantAccount)
	dependantGroup.Put("/:id",dependant.UpdateDependant)
	dependantGroup.Delete("/:id",dependant.DeleteDependantHandler)
}