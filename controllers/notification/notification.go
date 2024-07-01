package notification

import (
	"github.com/gofiber/fiber/v2"
	"main.go/model"
	"main.go/utilities"
)

func AddNotification(c *fiber.Ctx) error{
	body:=model.Notification{}
	if err := c.BodyParser(&body); err != nil {
		return utilities.ShowError(c,"failed to parse json data", fiber.StatusInternalServerError)
	}
	notification,err := model.SendNotification(c,body)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added notification",fiber.StatusOK,notification)
}

func GetNotification(c *fiber.Ctx)error{
	id,_:=model.GetAuthUserID(c)
	notification, err := model.GetNotification(c,id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved notification",fiber.StatusOK,notification)
}