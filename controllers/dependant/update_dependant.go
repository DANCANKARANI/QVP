package dependant

import (
	"github.com/gofiber/fiber/v2"
	"main.go/model"
	"main.go/utilities"
)

func UpdateDependant(c *fiber.Ctx)error{
	dependant_id :=c.Query("id")
	Dependant,err := model.UpdateDependant(c,dependant_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated dependant",fiber.StatusOK,Dependant)
}