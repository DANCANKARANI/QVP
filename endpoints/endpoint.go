package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/routes/dependants"
	"github.com/DANCANKARANI/QVP/routes/insurance"
	"github.com/DANCANKARANI/QVP/routes/notifications"
	"github.com/DANCANKARANI/QVP/routes/payment_methods"
	"github.com/DANCANKARANI/QVP/routes/payments"
	"github.com/DANCANKARANI/QVP/routes/users"
	"github.com/DANCANKARANI/QVP/routes/images"
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