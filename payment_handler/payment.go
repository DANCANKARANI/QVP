package payment_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/model"
	"main.go/utilities"
)

//adding payment handler
func AddPaymentHandler(c *fiber.Ctx)error{
	err := model.AddPayment(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"payments added successfully",fiber.StatusOK)
}

func UpdatePaymentHandler(c *fiber.Ctx)error{
	payment_id,_ :=uuid.Parse(c.Query("payment_id"))
	updated_payment,err :=model.UpdatePayment(c,payment_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated",fiber.StatusOK,updated_payment)
}

func GetPaymentsHandler(c *fiber.Ctx)error{
	payments,err:=model.GetPayments(c)
	if err!=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	var response []model.Payment
	for _,payment:=range payments{
		resPayment:=model.Payment{
			ID:payment.ID,
			Amount: payment.Amount,
			Narration: payment.Narration,
			Reference: payment.Reference,
			ResponseDescription: payment.ResponseDescription,
			UserID: payment.UserID,
			PaymentMethodID: payment.PaymentMethodID,
			User: payment.User,
			PaymentMethod: payment.PaymentMethod,
		}
		response = append(response, resPayment)
	}
	return utilities.ShowSuccess(c,"successfully retrieved the payments details",fiber.StatusOK,response)
}