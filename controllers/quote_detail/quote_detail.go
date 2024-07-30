package quote_detail

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add quote details handler
func AddQuoteDetailHandler(c *fiber.Ctx)error{
	response, err := model.AddQuoteDetail(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"quote details added successfully",fiber.StatusOK,response)
}

//update quote details handler
func UpdateQuoteDetailHandler(c *fiber.Ctx)error{
	id, _ := uuid.Parse(c.Params("id"))
	response,statusCode,err := model.UpdateQuoteDetail(c,id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),statusCode)
	}
	return utilities.ShowSuccess(c,"successfully updated quote details",statusCode,response)
}

//delete quote details handler
func DeleteQuoteDetailHandler(c *fiber.Ctx)error{
	id, _ := uuid.Parse(c.Params("id"))
	statusCode, err := model.DeleteQuoteDetail(c,id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),statusCode)
	}
	return utilities.ShowMessage(c,"quote details deleted successfully",statusCode)
}

//get quote details with prescription
func GetQuoteDetailWithPrescriptionHandler(c *fiber.Ctx)error{
	id,_ := uuid.Parse(c.Params("id"))
	statusCode,response, err := model.GetQuoteDetailsWithPrescription(id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),statusCode)
	}
	return utilities.ShowSuccess(c,"successfully retrieved quote details",statusCode,response)
}