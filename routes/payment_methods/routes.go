package payment_methods

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/controllers/payment"
	"github.com/DANCANKARANI/QVP/controllers/user"
)

func SetPaymentMethodRoutes(app *fiber.App){
	auth := app.Group("/api/v1/payment-methods")
	paymentMethodGroup := auth.Group("/",user.JWTMiddleware)
	paymentMethodGroup.Get("/",payment.GetPaymentMethodPaymentsHandler)
	paymentMethodGroup.Post("/",payment.AddPaymentMethod)
	paymentMethodGroup.Patch("/:id",payment.UpdatePaymentMethod)
	paymentMethodGroup.Delete("/:id",payment.RemovePaymentMethod)
}