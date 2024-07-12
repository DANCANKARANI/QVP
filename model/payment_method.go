package model

import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
adds the payments method
@params payment_method_id
@params payment_method PaymentMethod
*/
func AddPaymentMethod(c *fiber.Ctx,payment_method_id uuid.UUID, payment_method PaymentMethod) (*PaymentMethod, error) {
	paymentMethod := PaymentMethod{
		ID: payment_method_id,
		Title: payment_method.Title,
	}
	err:= db.Create(&paymentMethod).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to insert data")
	}
	return &paymentMethod,nil
}
/*
updates payment method details
@params c *fiber.Ctx
@payment_method_id
*/
func UpdtatePaymentMethod(c *fiber.Ctx,	payment_method_id string) (*PaymentMethod, error) {
	body := PaymentMethod{}
	if err := c.BodyParser(&body);err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to parse json data")
	}
	err := db.First(&PaymentMethod{},"id = ?",payment_method_id).Updates(&body).Scan(&body).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to update payment_method")
	}
	return &body, nil
}

/*
deletes the payment method
@params c *fiber.Ctx
@params payment_method_id
*/

func DeletePaymentMethod(c *fiber.Ctx,	payment_method_id string) error {
	payment_method :=PaymentMethod{}
	err:= db.First(&payment_method,"id = ?", payment_method_id).Delete(&payment_method).Error
	if err !=nil {
		log.Fatal(err.Error())
		return	errors.New("failed to delete payment method")
	}
	return nil
}

/*
returns the payments all methods
@params c*fiber.Ctx
*/
func GetPaymentMethods(c *fiber.Ctx) (*[]PaymentMethod,error) {
	var response []PaymentMethod

	// Check if there are any records found
	if err := db.Model(&PaymentMethod{}).Scan(&response).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil,errors.New("no payment methods found")
		}
		log.Fatalln(err.Error()+"jdsfss")
		return nil,errors.New("error getting payment methods")
	}

	// If no records are found, response will be empty but no error occurs
	if len(response) == 0 {
		return nil,errors.New("no payment methods found")
	}

	return &response,nil
}
