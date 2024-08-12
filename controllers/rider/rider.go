package rider

import (
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//gets the rider by id
func GetRiderHandler(c *fiber.Ctx)error{
id,_:=model.GetAuthUserID(c)
response,err:=model.GetRider(id)
if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved rider",fiber.StatusOK,response)
}

//update rider handler
func UpdateRiderHandler(c *fiber.Ctx)error{
	rider_id,_:=uuid.Parse(c.Params("id"))
	body:=model.Rider{}
	if err := c.BodyParser(&body);err != nil {
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to add rider",fiber.StatusForbidden)
	}
	IsValidData,err:= model.IsValidData(body.Email,body.PhoneNumber)
	if err != nil&& !IsValidData {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	response, err := model.UpdateRider(c,rider_id,body)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated rider",fiber.StatusOK,response)
}

//delete update handler
func DeleteRiderHandler(c *fiber.Ctx)error{
	rider_id,_:=uuid.Parse(c.Params("id"))
	err := model.DeleteRider(c,rider_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted rider",fiber.StatusOK)
}