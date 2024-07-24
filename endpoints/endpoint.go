package endpoints

import (
	"github.com/DANCANKARANI/QVP/routes/dependants"
	"github.com/DANCANKARANI/QVP/routes/images"
	"github.com/DANCANKARANI/QVP/routes/insurance"
	"github.com/DANCANKARANI/QVP/routes/modules"
	"github.com/DANCANKARANI/QVP/routes/notifications"
	"github.com/DANCANKARANI/QVP/routes/payment_methods"
	"github.com/DANCANKARANI/QVP/routes/payments"
	"github.com/DANCANKARANI/QVP/routes/permissions"
	"github.com/DANCANKARANI/QVP/routes/prescriptions"
	"github.com/DANCANKARANI/QVP/routes/riders"
	"github.com/DANCANKARANI/QVP/routes/roles"
	"github.com/DANCANKARANI/QVP/routes/teams"
	"github.com/DANCANKARANI/QVP/routes/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateEndpoint(){
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, change this to specific origins in production
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", 
	}))

	dependants.SetDependantRoutes(app)
	users.SetUserRoutes(app)
	payment_methods.SetPaymentMethodRoutes(app)
	payments.SetPaymentRoutes(app)
	insurance.SetInsuranceRoutes(app)
	notifications.SetNotificationRoute(app)
	images.SetImageRoutes(app)
	prescriptions.SetPrescriptionRoutes(app)
	riders.SetRiderRoutes(app)
	roles.SetRoleRoutes(app)
	permissions.SetPermissionRoutes(app)
	modules.SetModulesRoutes(app)
	teams.SetTeamRoutes(app)
	//port
	app.Listen(":3000")
}
