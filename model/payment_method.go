package model

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/utilities"
)

/*
adds the payments method
@params payment_method_id
@params payment_method PaymentMethod
*/
func AddPaymentMethod(c *fiber.Ctx,payment_method_id uuid.UUID, payment_method PaymentMethod) (*PaymentMethod, error) {
	IconURL,err:= utilities.GenerateImageUrl(c)
	if err != nil{
		return nil,err
	}
	paymentMethod := PaymentMethod{
		ID: payment_method_id,
		Title: payment_method.Title,
		IconUrl: IconURL,
	}
	err= db.Create(&paymentMethod).Error
	if err != nil {
		return nil, errors.New("failed to insert data:"+err.Error())
	}
	return &paymentMethod,nil
}
/*
updates payment method details
@params c *fiber.Ctx
*/
func UpdtatePaymentMethod(c *fiber.Ctx) (*PaymentMethod, error) {
	payment_method_id := c.Params("id")
	fmt.Println(payment_method_id)
	body := PaymentMethod{}
	if err := c.BodyParser(&body);err != nil {
		return nil,errors.New("failed to parse json data:"+err.Error())
	}
	err := db.First(&PaymentMethod{},"id = ?",payment_method_id).Updates(&body).Scan(&body).Error
	if err != nil {
		return nil, errors.New("failed to update payment_method:"+err.Error())
	}
	return &body, nil
}

/*
deletes the payment method
@params c *fiber.Ctx
*/

func DeletePaymentMethod(c *fiber.Ctx) error {
	payment_method_id := c.Params("id")
	payment_method :=PaymentMethod{}
	err:= db.First(&payment_method,"id = ?", payment_method_id).Delete(&payment_method).Error
	if err !=nil {
		return	errors.New("failed to delete payment method: "+ err.Error())
	}
	return nil
}

/*
returns the payments all methods
@params c*fiber.Ctx
*/
func GetPaymentMethods(c *fiber.Ctx)(*[]PaymentMethod,error){
	response:=[]PaymentMethod{}
	err:=db.Model(&PaymentMethod{}).Scan(&response).Error
	if err != nil{
		return nil,errors.New("Error getting payment methods:"+err.Error())
	}
	return &response,nil
}