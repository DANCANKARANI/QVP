package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
creates a sms
*/
func CreateSms(c *fiber.Ctx, callback_status string)(*Sms,error){
	sms := new(Sms)
	if err := c.BodyParser(&sms); err != nil{
		log.Println("failed to parse request data", err.Error())
		return nil, errors.New("failed to parse request data")
	}
	sms.ID=uuid.New()
	err := db.Create(&sms).Error
	if err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to add sms")
	}
	return sms, nil
}
/*
updates a sms
@params sms id
*/
func UpdateSms(c *fiber.Ctx, sms_id uuid.UUID,callback_status string)(*Sms, error){
	sms := new(Sms)
	if err := c.BodyParser(sms); err != nil{
		log.Println("failed to parse request data",err.Error())
		return nil, errors.New("failed to parse request data")
	}
	sms.CallbackStatus= callback_status
	err := db.Model(&sms).Where("id = ?",sms_id).Updates(&sms).Scan(&sms).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println("record not:",err.Error())
			return nil,errors.New("record not found")
		}
		log.Println("database error:",err.Error())
		return nil,errors.New("failed to update sms")
	}
	return sms, nil
}

/*
deletes sms
@params sms id
*/
func DeleteSms(sms_id uuid.UUID)error{
	sms := new(Sms)
	err := db.First(sms, "id = ?",sms_id).Delete(&sms).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println("record not found:", err.Error())
			return errors.New("record not found")
		}
		log.Println("database error:", err.Error())
		return errors.New("failed to delete sms")
	}
	return nil
}

/*
get sms by phone number
@params phone_number
*/
func GetUserSms(phone_number string)(*[]Sms,error){
	sms := new([]Sms)
	err := db.Model(sms).Where("phone = ?",phone_number).Find(&sms).Scan(&sms).Error
	if err != nil{
		log.Println("failed to get user's sms", err.Error())
		err_str := "failed to get sms for user with phone number:"+phone_number
		return nil,errors.New(err_str)
	}
	return sms, nil
}
