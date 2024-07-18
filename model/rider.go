package model

import (
	"errors"
	"log"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
var country_code = "KE"
func CreateRiderAccount(c *fiber.Ctx,body Rider)error {
	id:=uuid.New()
	body.ID = id
	if err :=db.Create(&body).Error;err !=nil {
		log.Println(err.Error())
		return errors.New("failed to add rider")
	}
	return nil;
}
//validate function
func IsValidData(body Rider)(bool,error){
	//validate email address
	_,err:= utilities.ValidateEmail(body.Email)
	if err != nil {
		return false,errors.New(err.Error())
	}
	//validate phone number
	_,err=utilities.ValidatePhoneNumber(body.PhoneNumber,country_code)
	if err != nil {
		return false,errors.New(err.Error())
	}
	return true,nil 
}

/*
gets the rider's details by id
@params id
*/
func GetRider(id uuid.UUID)(*Rider,error){
	rider:=Rider{}
	err:=db.First(&rider,"id = ?",id).Scan(&rider).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("no record found")
	}
	return &rider,nil
}