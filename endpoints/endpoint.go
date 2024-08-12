package endpoints

import (
	//"github.com/DANCANKARANI/QVP/routes/audits"
	"github.com/DANCANKARANI/QVP/routes/admins"
	"github.com/DANCANKARANI/QVP/routes/comments"
	"github.com/DANCANKARANI/QVP/routes/dependants"
	"github.com/DANCANKARANI/QVP/routes/images"
	"github.com/DANCANKARANI/QVP/routes/insurance"
	"github.com/DANCANKARANI/QVP/routes/insurance_users"
	"github.com/DANCANKARANI/QVP/routes/insurancers"
	"github.com/DANCANKARANI/QVP/routes/modules"
	"github.com/DANCANKARANI/QVP/routes/notifications"
	"github.com/DANCANKARANI/QVP/routes/payment_methods"
	"github.com/DANCANKARANI/QVP/routes/payments"
	"github.com/DANCANKARANI/QVP/routes/permissions"
	"github.com/DANCANKARANI/QVP/routes/prescriptions"
	"github.com/DANCANKARANI/QVP/routes/quote_details"
	"github.com/DANCANKARANI/QVP/routes/riders"
	"github.com/DANCANKARANI/QVP/routes/roles"
	"github.com/DANCANKARANI/QVP/routes/smss"
	"github.com/DANCANKARANI/QVP/routes/team_invitations"
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
	insurancers.SetInsurancersRoutes(app)
	admins.SetAdminsRoutes(app)
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
	team_invitations.SetTeamInvitationRoutes(app)
	insurance_users.SetInsuranceUserRoutes(app)
	quote_details.SetQuoteDetailsRoutes(app)
	comments.SetCommentRoutes(app)
	smss.SetSmsRoutes(app)

	
	//audits.SetAuditsRoutes(app)
	//port
	app.Listen(":3000")
}
