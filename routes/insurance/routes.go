package insurance

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/insurance"
	"main.go/controllers/user"
)

func SetInsuranceRoutes(app *fiber.App) {
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/insurance")
	insuranceGroup := auth.Group("/",user.JWTMiddleware)
	insuranceGroup.Get("/",insurance.GetAllInsuranceHandler)
	insuranceGroup.Get("/:id",insurance.GetOneInsuranceHandler)
	insuranceGroup.Post("/",insurance.AddInsuranceHandler )
	insuranceGroup.Put("/:id",insurance.UpdateInsuranceHandler)
}