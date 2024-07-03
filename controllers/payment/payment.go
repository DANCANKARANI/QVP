package payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
)
//adding payment handler
func AddPaymentHandler(c *fiber.Ctx)error{
	err := model.AddPayment(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"payments added successfully",fiber.StatusOK)
}

//get payments made by specific user
func GetUserPaymentsHandler(c *fiber.Ctx)error{
	response,err:=model.GetUserPayments(c)
	if err!=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved the payments details",fiber.StatusOK,response)
}

//get payments made by specific payment_method
func GetPaymentMethodPaymentsHandler(c *fiber.Ctx)error{
	response,err:=model.GetPaymentMethodPayments(c)
	if err!=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved the payments details",fiber.StatusOK,response)
}


//getting all payments
func GetAllPaymentsHandler(c *fiber.Ctx)error{
	user_id,_:=model.GetAuthUserID(c)
	payments,err := model.GetAllPayments(c,user_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	
	return utilities.ShowSuccess(c,"successfully retrieved payments",fiber.StatusOK,payments)
}

//update payment handler
func UpdatePaymentHandler(c *fiber.Ctx)error{
	payment,err:=model.UpdatePayment(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated payments",fiber.StatusOK,payment)
}

//deletes a row of a payment
func DeletePaymentHandler(c *fiber.Ctx)error{
	err:=model.DeletePayment(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"record deleted",fiber.StatusOK)
}