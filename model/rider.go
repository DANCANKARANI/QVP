package model

import (
	"errors"
	"log"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
var country_code = "KE"

//response to rider
type ResRider struct{
	ID 		uuid.UUID		`json:"id"`
	FullName	string		`json:"full_name"`
	StaffMember	string		`json:"staff_member"`
	PhoneNumber	string		`json:"phone_number"`
	Email		string		`json:"email"`
	IdNumber	string		`json:"id_number"`
}
func CreateRiderAccount(c *fiber.Ctx,body Rider)(uuid.UUID,error ){
	id:=uuid.New()
	body.ID = id
	if err :=db.Create(&body).Error;err !=nil {
		log.Println(err.Error())
		return uuid.Nil,errors.New("failed to add rider")
	}
	return body.ID,nil
}
//validate function
func IsValidData(email, phone_number string)(bool,error){
	//validate email address
	_,err:= utilities.ValidateEmail(email)
	if err != nil {
		return false,errors.New(err.Error())
	}
	//validate phone number
	_,err=utilities.ValidatePhoneNumber(phone_number,country_code)
	if err != nil {
		return false,errors.New(err.Error())
	}
	return true,nil 
}

/*
gets the rider's details by id
@params id
*/
func GetRider(id uuid.UUID)(*ResRider,error){
	rider:=Rider{}
	response := new(ResRider)
	err:=db.First(&rider,"id = ?",id).Scan(response).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("no record found")
	}
	return response,nil
}
/*
updates rider
@params rider_id
*/
func UpdateRider(c *fiber.Ctx, rider_id uuid.UUID,body Rider)(*ResRider,error){
	user_id, _:=GetAuthUserID(c)
	role := GetAuthUser(c)

	//parse request body

	response := ResRider{}
	oldValues := new(Rider)
	//find rider
	err := db.First(oldValues,"id = ?",rider_id).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to update rider")
	}
	//update rider
	err = db.Model(&oldValues).Updates(&body).Scan(&response).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to update rider")
	}
	newValues := oldValues

	//update audit logs
	if err := utilities.LogAudit("Update",user_id,role,"Rider",rider_id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return &response,nil
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

//check if rider already exists
