package model

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/utilities"
)

func AddPayment(c *fiber.Ctx) error{
	id := uuid.New()
	user_id,err := GetAuthUserID(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}

	body := Payment{}
	if err :=c.BodyParser(&body); err != nil {
		return errors.New("failed to parse body JSON "+err.Error())
	}
	payment_method_id,_ := uuid.Parse(c.Query("payment_method_id"))
	//initializing the id's
	body.ID=id
	body.UserID=user_id
	body.PaymentMethodID=payment_method_id
	result := db.Create(&body)
	if result.Error != nil {
		return errors.New("failed to add payment:"+result.Error.Error())
	}
	return nil
}

/*
updates the payment details
@params payment_id
*/

func UpdatePayment(c *fiber.Ctx,payment_id uuid.UUID)(*Payment,error){
	body := Payment{}
	if err := c.BodyParser(&body); err != nil{
		return &Payment{}, errors.New("failed to parse json data")
	}
	result := db.Model(&body).Where("id = ?",payment_id).Updates(&body)
	if result.Error != nil {
		return &Payment{}, errors.New("failed to update payments: " + result.Error.Error())
	}
	return &body,nil
}

/*
gets the payments
*/

func GetPayments(c *fiber.Ctx)([]Payment,error){
	payments :=[]Payment{}
	result := db.Model(&Payment{}).Preload("users").Preload("payment_methods").Find(&payments).Scan(payments)
	if result.Error != nil {
		return payments,errors.New(result.Error.Error())
	}
	return payments,nil
}