package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
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
	user_id, _:=GetAuthUserID(c)

	role := GetAuthUser(c)

	paymentMethod := PaymentMethod{
		ID: payment_method_id,
		Title: payment_method.Title,
	}
	err:= db.Create(&paymentMethod).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to insert data")
	}
	newValues := paymentMethod

	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"Payment method",payment_method_id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return &paymentMethod,nil
}
/*
updates payment method details
@params c *fiber.Ctx
@payment_method_id
*/
func UpdtatePaymentMethod(c *fiber.Ctx,	payment_method_id uuid.UUID) (*PaymentMethod, error) {
	body := PaymentMethod{}

	user_id, _ := GetAuthUserID(c)

	role := GetAuthUser(c)

	//get request body
	paymentMethod := new(PaymentMethod)
	if err := c.BodyParser(&body);err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to parse json data")
	}
	//find record and old values
	err := db.First(&paymentMethod,"id = ?",payment_method_id).Error
	if err != nil{
		log.Println("error finding payment method",err.Error())
		return nil, errors.New("failed to update payment method")
	}

	oldValues := paymentMethod

	//update payment record
	if err := db.Model(&paymentMethod).Updates(&body).Error; err != nil{
		log.Println("error updating payment method")
		return nil, errors.New("failed to update payment method")
	}

	newValues := paymentMethod

	//update audit logs
	if err := utilities.LogAudit("Update",user_id,role,"Payment method",payment_method_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	//response
	return newValues, nil
}

/*
deletes the payment method
@params c *fiber.Ctx
@params payment_method_id
*/
func DeletePaymentMethod(c *fiber.Ctx,	payment_method_id uuid.UUID) error {
	user_id, _:= GetAuthUserID(c)

	role := GetAuthUser(c)

	payment_method :=PaymentMethod{}

	//find payment_method record
	err:= db.First(&payment_method,"id = ?", payment_method_id).Error
	if err != nil{
		log.Println("error finding payment_method for deleting",err.Error())
		return errors.New("failed to delete payment_method")
	}
	oldValues := payment_method
	
	//delete payment method
	err = db.Delete(&payment_method).Error
	if err !=nil {
		log.Fatal(err.Error())
		return	errors.New("failed to delete payment method")
	}
	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"Payment method",payment_method_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
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
