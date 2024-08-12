package insurancers

import (
	"github.com/DANCANKARANI/QVP/controllers/insurancer"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetInsurancersRoutes(app *fiber.App) {
	// Group routes under /api/v1/insurancers
	auth := app.Group("/api/v1/insurancers")
	auth.Post("/1/signup",insurancer.CreateInsurancerAccountHandler)
	auth.Post("/login",insurancer.InsurancerLogin)

	insurancerGroup := auth.Group("/", user.JWTMiddleware)
	insurancerGroup.Get("/",insurancer.GetInsurancerHandler)
	insurancerGroup.Post("/logout",insurancer.Logout)
	insurancerGroup.Patch("/",insurancer.UpdateInsurancer)
	insurancerGroup.Delete("/",insurancer.DeleteInsurancerHandler)
}