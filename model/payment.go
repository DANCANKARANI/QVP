package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
type ResponsePayment struct {
	ID uuid.UUID `json:"id"`
	Amount float64 `json:"amount"`
	Narration string `json:"narration"`
	Reference string `json:"reference"`
	ResponseDescription string `json:"response_description"`
	UserID uuid.UUID `json:"user_id"`
	PaymentMethodId uuid.UUID `json:"payment_method_id"`
	User ResponseUser
	PaymentMethod PaymentMethod
}
func AddPayment(c *fiber.Ctx,user_id,payment_method_id uuid.UUID) error{
	role := GetAuthUser(c)
	//generate new uuid
	id := uuid.New()
	body := Payment{}
	if err :=c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return errors.New("failed to parse payments data")
	}

	//initializing the id's
	body.ID=id
	body.UserID=user_id
	body.PaymentMethodID=payment_method_id
	err := db.Create(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to add payment")
	}

	newValues := body

	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"Payment",id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return nil
}



/*
gets the payments for a specific user
@params user_id
*/

func GetUserPayments(c *fiber.Ctx,user_id uuid.UUID)(*[]ResponsePayment,error){
	var payments []Payment
    if err := db.Preload("User").Preload("PaymentMethod").
        Where("user_id = ?", user_id).Find(&payments).Error; err != nil {
        return nil,errors.New("failed to get data:"+err.Error())
    }
	var response []ResponsePayment
	for _,payment := range payments{
		resPaymentMethod :=PaymentMethod{
			ID: payment.PaymentMethod.ID,
			Title: payment.PaymentMethod.Title,
			ImageID: payment.PaymentMethod.ImageID,
		}
		resUser := ResponseUser{
			ID: payment.User.ID,
			FullName: payment.User.FullName,
			PhoneNumber: payment.User.PhoneNumber,
			Email: payment.User.Email,
		}
		resPayment :=  ResponsePayment{
			ID: payment.ID,
			Amount: payment.Amount,
			Narration: payment.Narration,
			Reference: payment.Reference,
			ResponseDescription: payment.ResponseDescription,
			PaymentMethodId: payment.PaymentMethodID,
			UserID: payment.UserID,
			User: resUser,
			PaymentMethod: resPaymentMethod,
		}
		response = append(response,resPayment)
	}
	return &response, nil
}

//get all payments

func GetAllPayments(c *fiber.Ctx)(*[]ResponsePayment,error){
	var payments []Payment
    if err := db.Preload("User").Preload("PaymentMethod").
       // Where("user_id = ?", user_id)
	    Find(&payments).Error; err != nil {
			log.Println(err.Error())
        return nil,errors.New("failed to get data")
    }
	var response []ResponsePayment
	for _,payment := range payments{
		resPaymentMethod :=PaymentMethod{
			ID: payment.PaymentMethod.ID,
			Title: payment.PaymentMethod.Title,
			ImageID: payment.PaymentMethod.ImageID,
		}
		resUser := ResponseUser{
			ID: payment.User.ID,
			FullName: payment.User.FullName,
			PhoneNumber: payment.User.PhoneNumber,
			Email: payment.User.Email,
		}
		resPayment :=  ResponsePayment{
			ID: payment.ID,
			Amount: payment.Amount,
			Narration: payment.Narration,
			Reference: payment.Reference,
			ResponseDescription: payment.ResponseDescription,
			PaymentMethodId: payment.PaymentMethodID,
			UserID: payment.UserID,
			User: resUser,
			PaymentMethod: resPaymentMethod,
		}
		response = append(response,resPayment)
	}
	return &response, nil
}

/*
updates the payment details
*/
func UpdatePayment(c *fiber.Ctx)(*ResponsePayment,error){
	user_id, _ := GetAuthUserID(c)

	role := GetAuthUser(c)

	payment_id,_ :=uuid.Parse(c.Params("id"))

	body := new (Payment)
	payment := new(Payment)
	response := new(ResponsePayment)

	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to parse payment data")
	}
	//get and update the payment
	if err := db.First(&payment,"id = ?",payment_id).Error; err != nil{
		log.Println("eror updating payment", err.Error())
		return nil, errors.New("failed to update payment")
	}

	oldValues := payment
	
	//update payment
	err := db.Preload("User").Preload("PaymentMethod").
	Model(payment).Updates(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update payment")
	}
	newValues := payment

	//update log audits
	if err := utilities.LogAudit("Update",user_id,role,"Payment",payment_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return response,nil
}
/*
deletes a specified raw
@params payment_id
*/
func DeletePayment(c *fiber.Ctx,payment_id uuid.UUID) error {
	user_id, _ := GetAuthUserID(c)

	role := GetAuthUser(c)
	// Perform the delete operation
	payment := new(Payment)
	err := db.First(&payment,"id = ?",payment_id).Error
	if err != nil {
		log.Println("error deleting payment",err.Error())
		return errors.New("failed to delete payment")
	}
	oldValues := payment
	
	err =db.Delete(&payment).Error
	if err != nil {
		log.Println("error deleting payment",err.Error())
		return errors.New("failed to delete the record")
	}

	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"Payment",payment_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
	return nil
}