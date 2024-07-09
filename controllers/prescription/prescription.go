package prescription

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add prescription handler
func AddPrescriptionHandler(c *fiber.Ctx)error{
	//admin_id,_:= model.GetAuthUserID(c)
	user_id,_ :=uuid.Parse(c.Query("user_id"))
	rider_id,_ :=uuid.Parse(c.Query("rider_id"))
	prescription,err := model.AddPrescription(c,user_id,rider_id,uuid.Nil)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"Prescription added successfully",fiber.StatusOK,prescription)
}
//get prescrption handler
func GetPrescriptionsHandler(c *fiber.Ctx)error{
	//id,_ := model.GetAuthUserID(c)
	id := c.Query("user_id")
	response,err:=model.GetUsersPrescription(c,id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"prescription retrieved successfully",fiber.StatusOK,response)
}
//update prescription handler
func UpdatePrescriptionHandler(c *fiber.Ctx)error{
	id,_:= uuid.Parse(c.Params("id"))
	admin_id,_:= model.GetAuthUserID(c)
	user_id,_ :=uuid.Parse(c.Params("user_id"))
	rider_id,_ :=uuid.Parse(c.Params("rider_id"))
	response,err := model.UpdatePrescription(c,id,user_id,rider_id,admin_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated prescription",fiber.StatusOK,response)
}
//Delete prescription handler
func DeletePrescriptionHandler(c *fiber.Ctx)error{
	id := c.Params("id")
	err := model.DeletePrescription(c,id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"prescription deleted successfully",fiber.StatusOK)
}