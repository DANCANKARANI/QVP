package notifications

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/controllers/notification"
	"github.com/DANCANKARANI/QVP/controllers/user"
)

func SetNotificationRoute(app *fiber.App){
	auth := app.Group("/api/v1/notifications")
	notificationGroup := auth.Group("/",user.JWTMiddleware)
	notificationGroup.Post("/", notification.AddNotification)
	notificationGroup.Get("/",notification.GetNotification)
}