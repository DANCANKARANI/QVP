package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/dependant_handler"
	"main.go/insurance_handler"
	"main.go/payment_handler"
	"main.go/user_handler"
)

func EndPoints() {
	app := fiber.New()
	app.Post("/register",user_handler.CreateUserAccount)
	app.Post("/login",user_handler.Login)
	app.Post("/forgot-password",user_handler.ForgotPassword)
	app.Post("/reset-password",user_handler.ResetPassword)
	//start the server

	//JWT middleware
	app.Use(user_handler.JWTMiddleware)
	
	//protected routes

	//users routes
	app.Get("/get/one/user",user_handler.GetOneUserHandler)
	app.Get("/get/all/users",user_handler.GetAllUsersHandler)
	app.Patch("/update/user",user_handler.UpdateUserHandler)
	//dependants routes
	app.Post("/register-dependant", dependant_handler.RegisterDependantAccount)
	app.Get("/get-dependants",dependant_handler.GetDependantsHandler)
	app.Get("/get-dependant-id",dependant_handler.GetDependantID)
	app.Patch("/update/dependants",dependant_handler.UpdateDependant)
	app.Delete("/delete/dependant",dependant_handler.DeleteDependantHandler)

	//payment_methods routes
	app.Post("/add/payment-method",payment_handler.AddPaymentMethod)
	app.Patch("/update/payment-method",payment_handler.UpdatePaymentMethod)
	app.Get("/get/payment-method",payment_handler.GetPaymentMethods)
	app.Delete("/delete/payment-method",payment_handler.RemovePaymentMethod)
	//payments routes
	app.Post("/add/payment", payment_handler.AddPaymentHandler)
	app.Get("/get/user/payments",payment_handler.GetUserPaymentsHandler)
	app.Get("/get/payment-method/payments",payment_handler.GetPaymentMethodPaymentsHandler)
	app.Get("/get/all/payments",payment_handler.GetAllPaymentsHandler)
	app.Patch("/update/payments",payment_handler.UpdatePaymentHandler)
	app.Delete("/delete/payment",payment_handler.DeletePaymentHandler)
	//Insurance routes
	app.Post("/add/insurance",insurance_handler.AddInsuranceHandler)
	app.Post("/update/insurance",insurance_handler.UpdateInsuranceHandler)
	app.Get("/get/one/insurance",insurance_handler.GetOneInsuranceHandler)
	app.Get("/get/all/insurances",insurance_handler.GetAllInsuranceHandler)
	app.Get("/logout",user_handler.Logout)

	app.Listen(":3000")
}
