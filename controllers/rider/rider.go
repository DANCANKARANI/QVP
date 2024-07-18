package rider

import (
	"log"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)
//creates rider's account
func RegisterRider(c *fiber.Ctx) error {
	body:=model.Rider{}
	if err := c.BodyParser(&body);err != nil {
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to add rider",fiber.StatusForbidden)
	}
	IsValidData,err:= model.IsValidData(body)
	if err != nil&& !IsValidData {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	err = model.CreateRiderAccount(c,body)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"rider created successfully",fiber.StatusOK)
}
//gets the rider by id
func GetRiderHandler(c *fiber.Ctx)error{
id,_:=model.GetAuthUserID(c)
response,err:=model.GetRider(id)
if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved rider",fiber.StatusOK,response)
}