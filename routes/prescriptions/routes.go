package prescriptions

import (
	"github.com/DANCANKARANI/QVP/controllers/prescription"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetPrescriptionRoutes(app *fiber.App) {

	auth := app.Group("/api/v1/prescription")
	prescriptionGroup := auth.Group("/",user.JWTMiddleware)
	prescriptionGroup.Post("/",prescription.AddPrescriptionHandler)
	prescriptionGroup.Get("/",prescription.GetPrescriptionsHandler)
	prescriptionGroup.Patch("/:id",prescription.UpdatePrescriptionHandler)
	prescriptionGroup.Delete("/:id",prescription.DeletePrescriptionHandler)
}