package insurance_user

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//assigns insurance users
func AssignInsuranceUserHandler(c *fiber.Ctx) error {
	response,err := model.AssisgnInsuranceUser(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully assigned insurance users",fiber.StatusOK,response)
}

//updates insurance users
func UpdateInsuranceUserHandler(c *fiber.Ctx)error{
	id,_ := uuid.Parse(c.Params("id"))
	response,err := model.UpdateInsuranceUser(c,id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated insurance user",fiber.StatusOK,response)
}

//deletes insurance user
func DeleteInsuranceUserHandler(c *fiber.Ctx)error{
	id,_ := uuid.Parse(c.Params("id"))
	err := model.DeleteInsuranceUser(id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted insurance user",fiber.StatusOK)
}
//gets user with insurance
func GetUserWithInsuranceHandler(c *fiber.Ctx)error{
	//id,_ := uuid.Parse(c.Params("id"))
	response, err := model.GetUsersWithInsurance()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved users with insurance",fiber.StatusOK,response)
}