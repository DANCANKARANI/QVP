package dependant_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/model"
	"main.go/utilities"
)

func UpdateDependant(c *fiber.Ctx)error{
	dependant_id :=c.Params("id")
	dependantUUID, err := uuid.Parse(dependant_id)
	if err != nil {
		return utilities.ShowError(c,"failed to convert id string to uuid",fiber.StatusInternalServerError)
	}
	Dependant,err := model.UpdateDependant(c,dependantUUID)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated dependant",fiber.StatusOK,Dependant)
}