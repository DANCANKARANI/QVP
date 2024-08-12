package rider

import (
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
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
	rider_id,_:=model.GetAuthUserID(c)
	body:=model.Rider{}
	if err := c.BodyParser(&body);err != nil {
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to add rider",fiber.StatusForbidden)
	}

	if body.PhoneNumber !="" && body.Email !=""{
			IsValidData,err:= model.IsValidData(body.Email,body.PhoneNumber)
			if err != nil&& !IsValidData {
			return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
		}
	}

	//hash password
	if body.Password != ""{
		hashed_password,_ := utilities.HashPassword(body.Password)
		body.Password = hashed_password
	}


	response, err := model.UpdateRider(c,rider_id,body)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated rider",fiber.StatusOK,response)
}

//delete update handler
func DeleteRiderHandler(c *fiber.Ctx)error{
	rider_id,_:=model.GetAuthUserID(c)
	err := model.DeleteRider(c,rider_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted rider",fiber.StatusOK)
}