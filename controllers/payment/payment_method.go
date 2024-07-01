package payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/model"
	"main.go/utilities"
)
var body = model.PaymentMethod{}
//add payment handler
func AddPaymentMethod(c *fiber.Ctx) error {
	payment_method_id := uuid.New()
	if err := c.BodyParser(&body); err != nil {
		return utilities.ShowError(c,"faied to parse json data",fiber.StatusInternalServerError)
	}
	payment_method,err := model.AddPaymentMethod(c,payment_method_id,body)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"succesfully added payment method",fiber.StatusOK,payment_method)
}

//update payment handler
func UpdatePaymentMethod(c *fiber.Ctx)error{
	
	payment_method,err := model.UpdtatePaymentMethod(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(), fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated payment method",fiber.StatusOK,payment_method)
}

//deleting the payment method
func RemovePaymentMethod(c *fiber.Ctx)error{
	err :=model.DeletePaymentMethod(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"payment method removed successfully",fiber.StatusOK)
}

//getting the payment method handler
func GetPaymentMethods(c *fiber.Ctx)error{
	response,err := model.GetPaymentMethods(c)
	if err != nil {
		return utilities.ShowError(c,"error getting the payment methods",fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrived payment methods", fiber.StatusOK,response)
}