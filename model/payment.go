package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	payment_id :=c.Params("id")
	body := Payment{}
	response := ResponsePayment{}
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to parse payment data")
	}
	//get and update the payment
	err := db.Preload("User").Preload("PaymentMethod").
	First(&Payment{},"id = ?",payment_id).Updates(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update payment")
	}
	return &response,nil
}
/*
deletes a specified raw
@params payment_id
*/
func DeletePayment(c *fiber.Ctx,payment_id string) error {
	// Perform the delete operation
	err := db.First(&Payment{},"id = ?",payment_id).Delete(&Payment{}).Error
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Payment record not found")
		}
		return errors.New("failed to delete the record")
	}
	return nil
}