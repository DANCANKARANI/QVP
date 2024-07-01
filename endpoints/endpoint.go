package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"main.go/routes/dependants"
	"main.go/routes/insurance"
	"main.go/routes/notifications"
	"main.go/routes/payment_methods"
	"main.go/routes/payments"
	"main.go/routes/users"
	"main.go/routes/images"
)

func CreateEndpoint(){
	app := fiber.New()
	dependants.SetDependantRoutes(app)
	users.SetUserRoutes(app)
	payment_methods.SetPaymentMethodRoutes(app)
	payments.SetPaymentRoutes(app)
	insurance.SetInsuranceRoutes(app)
	notifications.SetNotificationRoute(app)
	images.SetImageRoutes(app)
	app.Listen(":3000")
}