package dependant

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateDependant(c *fiber.Ctx)error{
	dependant_id,_ := uuid.Parse(c.Params("id"))
	Dependant,err := model.UpdateDependant(c,dependant_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated dependant",fiber.StatusOK,Dependant)
}