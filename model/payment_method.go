package model

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
adds the payments method
@params payment_method_id
@params payment_method PaymentMethod
*/
func AddPaymentMethod(c *fiber.Ctx,payment_method_id uuid.UUID, payment_method PaymentMethod) (PaymentMethod, error) {
	paymentMethod := PaymentMethod{
		ID: payment_method_id,
		Title: payment_method.Title,
		IconUrl: payment_method.IconUrl,
	}
	result := db.Create(&paymentMethod)
	if result.Error == nil {
		return paymentMethod, result.Error
	}
	return paymentMethod,nil
}
/*
updates payment method details
@params c *fiber.Ctx
@params payment_method_id
@params payment_method PaymentMethod
*/
func UpdtatePaymentMethod(c *fiber.Ctx,payment_method_id string,payment_method PaymentMethod) (PaymentMethod, error) {
	result := db.Model(&PaymentMethod{}).Where("id = ?",payment_method_id).Updates(&payment_method)
	if result.Error != nil {
		return PaymentMethod{}, errors.New(result.Error.Error()+" update")
	}
	return payment_method,nil
}

/*
deletes the payment method
@params payment_method_id
*/

func DeletePaymentMethod(c *fiber.Ctx, payment_method_id string) error {
	result := db.Model(&PaymentMethod{}).Where("id = ?", payment_method_id)
	if result.Error !=nil {
		return result.Error
	}
	db.Delete(&PaymentMethod{})
	return nil
}

/*
returns the payments all methods
*/

func GetPaymentMethods(c *fiber.Ctx)([]PaymentMethod,error){
	PaymentMethods:=[]PaymentMethod{}
	result:=db.Find(&PaymentMethods).Scan(&PaymentMethods)
	if result.Error != nil{
		return PaymentMethods,result.Error
	}
	return PaymentMethods,nil
}