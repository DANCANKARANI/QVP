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
/*
updates rider
@params rider_id
*/
func UpdateRider(c *fiber.Ctx, rider_id uuid.UUID, body Rider)(*Rider,error){
	user_id, _:=GetAuthUserID(c)
	role := GetAuthUser(c)

	oldValues := body
	//find rider
	err := db.First(oldValues,"id = ?",rider_id).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to update rider")
	}
	//update rider
	newValues := new(Rider)
	err = db.Model(&oldValues).Where("id = ?",rider_id).Scan(&newValues).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to update rider")
	}

	//update audit logs
	if err := utilities.LogAudit("Update",user_id,role,"Rider",rider_id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return newValues,nil
}

/*
deletes rider
@params rider_id
*/
func DeleteRider(c *fiber.Ctx, rider_id uuid.UUID)(error){
	user_id, _ :=GetAuthUserID(c)
	role := GetAuthUser(c)

	oldValues := new(Rider)
	err := db.First(&oldValues,"id = ?",rider_id).Delete(&oldValues).Error
	if err != nil {
		log.Println(err.Error())
		return  errors.New("failed to delete rider")
	}
	//update log audit
	if err := utilities.LogAudit("Delete",user_id,role,"Rider",rider_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
	//response
	return nil
}