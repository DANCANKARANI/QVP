package payment

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//adding payment handler
func AddPaymentHandler(c *fiber.Ctx)error{
	user_id,_:=model.GetAuthUserID(c)
	payment_method_id,_ := uuid.Parse(c.Query("payment_method_id"))
	err := model.AddPayment(c,user_id,payment_method_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"payments added successfully",fiber.StatusOK)
}

//get payments made by a specific user
func GetUserPaymentsHandler(c *fiber.Ctx)error{
	user_id,_:=model.GetAuthUserID(c)
	payments,err := model.GetUserPayments(c,user_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	
	return utilities.ShowSuccess(c,"successfully retrieved payments",fiber.StatusOK,payments)
}


//getting all payments
func GetAllPaymentsHandler(c *fiber.Ctx)error{
	payments,err := model.GetAllPayments(c)
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
	payment_id := c.Params("id")
	err:=model.DeletePayment(c,payment_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"record deleted",fiber.StatusOK)
}