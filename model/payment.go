package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/utilities"
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
gets the payments for a specific user
*/

func GetUserPayments(c *fiber.Ctx)(*[]ResponsePayment,error){
	user_id,err:=GetAuthUserID(c)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	response :=[]ResponsePayment{}
	 err = db.Model(&Payment{}).Where("user_id = ?", user_id).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to get data:"+err.Error())
	}
	return &response,nil
}

//Get the payments made buy a specific payment method
func GetPaymentMethodPayments(c *fiber.Ctx)(*[]ResponsePayment,error){
	payment_method_id,_ := uuid.Parse(c.Query("id"))
	response := []ResponsePayment{}
	err := db.Model(&Payment{}).Where("payment_method_id = ?", payment_method_id).Scan(&response).Error
	if err !=nil{
		return nil, err
	}
	return &response, nil
}

//get all payments

func GetAllPayments(c *fiber.Ctx,user_id uuid.UUID)(*[]ResponsePayment,error){
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
			IconUrl: payment.PaymentMethod.IconUrl,
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
	payment_id,err:=uuid.Parse(c.Query("id"))
	if err != nil {
		return nil,err
	}
	body := Payment{}
	response := ResponsePayment{}
	if err := c.BodyParser(&body); err != nil {
		return nil,errors.New("failed to parse json data:"+err.Error())
	}
	err = db.First(&Payment{},"id = ?",payment_id).Updates(&body).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to update payment:"+err.Error())
	}
	return &response,nil
}
/*
deletes a specified raw
*/
func DeletePayment(c *fiber.Ctx)error{
	payment_id,err:=uuid.Parse(c.Query("id"))
	if err != nil {
		return err
	}
	err =db.Model(&Payment{}).Where("id = ?",payment_id).Delete(&Payment{}).Error
	if err != nil {
		return errors.New("failed to delete the record:"+err.Error())
	}
	return nil
}