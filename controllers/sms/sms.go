package sms

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add sms handler
func AddSmsHandler(c *fiber.Ctx)error{
	response, err := model.CreateComment(c)
	if err != nil{
		return utilities.ShowError(c, err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"sms added successfully",fiber.StatusOK,response)
}

//update sms handler
func UpdateSmsHandler(c *fiber.Ctx)error{
	sms_id, _ := uuid.Parse(c.Params("id"))
	status := c.Query("sent")
	response, err := model.UpdateSms(c,sms_id,status)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"sms updated successfully",fiber.StatusOK,response)
}

//delete sms handler
func DeleteSmsHandler(c *fiber.Ctx)error{
	sms_id,_ := uuid.Parse(c.Params("id"))
	err := model.DeleteSms(sms_id)
	if err != nil{
		return utilities.ShowError(c, err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully",fiber.StatusOK)
}
//get users' sms handler
func GetUserSmsHandler(c *fiber.Ctx)error{
	phone_number:=c.Params("phone")
	response, err := model.GetUserSms(phone_number)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved users sms",fiber.StatusOK,response)
}