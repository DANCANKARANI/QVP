package payments

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/payment"
	"main.go/controllers/user"
)

func SetPaymentRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/payments")
	paymentGroup := auth.Group("/",user.JWTMiddleware)
	paymentGroup.Get("/",payment.GetAllPaymentsHandler)
	paymentGroup.Post("/",payment.AddPaymentHandler)
	paymentGroup.Patch("/:id",payment.UpdatePaymentHandler)
	paymentGroup.Delete("/:id",payment.DeletePaymentHandler)
}