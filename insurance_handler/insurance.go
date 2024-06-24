package insurance_handler

import (
	"github.com/gofiber/fiber/v2"
	"main.go/model"
	"main.go/utilities"
)
//add insurance handler
func AddInsuranceHandler(c *fiber.Ctx)error{
	err := model.AddInsurace(c)
	if err !=nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c, "insurance added successfully",fiber.StatusOK)
}

//update insurance handler
func UpdateInsuranceHandler(c *fiber.Ctx)error{
	insurance,err := model.UpdateInsurance(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"insurance updated successfully",fiber.StatusOK,insurance)
}

//get insurance by id
func GetOneInsuranceHandler(c *fiber.Ctx)error{
	Insurance,err := model.GetOneInsurance(c)
	if err != nil {
		 return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"insurance retrieved successfully",fiber.StatusOK,Insurance)
}

//get all insurances
func GetAllInsuranceHandler(c *fiber.Ctx)error{
	insurances,err:=model.GetAllInsurances(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"insurances retrieved successfully",fiber.StatusOK,insurances)
}