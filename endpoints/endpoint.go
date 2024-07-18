package endpoints

import (
	"github.com/DANCANKARANI/QVP/routes/dependants"
	"github.com/DANCANKARANI/QVP/routes/images"
	"github.com/DANCANKARANI/QVP/routes/insurance"
	"github.com/DANCANKARANI/QVP/routes/notifications"
	"github.com/DANCANKARANI/QVP/routes/payment_methods"
	"github.com/DANCANKARANI/QVP/routes/payments"
	"github.com/DANCANKARANI/QVP/routes/prescriptions"
	"github.com/DANCANKARANI/QVP/routes/riders"
	"github.com/DANCANKARANI/QVP/routes/users"
	"github.com/gofiber/fiber/v2"
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
	prescriptions.SetPrescriptionRoutes(app)
	riders.SetRiderRoutes(app)
	app.Listen(":3000")
}