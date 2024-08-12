package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
creates insurancer account
@params body
*/
func CreateInsurancerAccount(c *fiber.Ctx,body Insurancer)error{
	Insurancer :=body

	//hash password
	hashed_password,err := utilities.HashPassword(Insurancer.Password)
	if err != nil{
		return errors.New(err.Error())
	}

	Insurancer.Password = hashed_password
	Insurancer.ID=uuid.New()
	//create insurancer
	if err := db.Create(&Insurancer).Error; err != nil{
		log.Println("error creating insurancer account:",err.Error())
		return errors.New("failed to create insurancer account")
	}

	return nil
}

/*
find insurer existence
*/
func InsurerExist(c *fiber.Ctx, phoneNumber string) (bool, *Insurancer, error) {
	var  existingUser Insurancer
 
	 // Detailed logging
	 log.Printf("Checking for user with phone number: %s", phoneNumber)
 
	 result := db.Where("phone_number = ?", phoneNumber).First(&existingUser)
	 if result.Error != nil {
		 // Log the detailed error
		 log.Printf("Error finding user with phone number %s: %v", phoneNumber, result.Error)
 
		 if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			 log.Println("record not found:",result.Error.Error())
			 return false, nil, nil
		 }
 
		return false, nil, fmt.Errorf("database error: %v", result.Error)
	}
	log.Printf("User found: %+v", existingUser)
	return true, &existingUser, nil
}

/*
update Insurancer
*/

func UpdateInsurancer(c *fiber.Ctx, insurancer_id uuid.UUID)(*Insurancer,error){
	insurancer := new(Insurancer)
	body := Insurancer{}

	//parse request
	if err :=  c.BodyParser(&body); err != nil{
		log.Println("error parsing request body:", err.Error())
		return nil, errors.New("error parsing request body")
	}
	if body.PhoneNumber != "" || body.Email != ""{
		valid,err :=IsValidData(body.Email,body.PhoneNumber)
		if !valid{
			return nil, err
		}

		//check if the phone number is already used
		exist,_,_:=InsurerExist(c,body.PhoneNumber)
		if exist{
			err_str := "phone number:"+body.PhoneNumber+" is already in use"
			return nil, errors.New(err_str)
		}
	}

	//find old values
	if err := db.First(insurancer,"id = ?",insurancer_id).Error; err != nil{
		log.Println("error getting insurancer for update:",err.Error())
		return nil, errors.New("failed to update insurancer")
	}
	oldValues := insurancer

	//update insurancer
	if err := db.Model(&insurancer).Updates(body).Error; err != nil{
		log.Println("failed to update insurancer:",err.Error())
		return nil,errors.New("failed to update insurancer")
	}
	newValues := insurancer

	user_id,_ := GetAuthUserID(c)
	role := GetAuthUser(c)
	//update audit logs
	if err := utilities.LogAudit("Update",user_id,role,"Insurancer",insurancer_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return newValues, nil
}

/*
delete insurancer
@params insurancer_id
*/
func DeleteInsurancer(c *fiber.Ctx, insurancer_id uuid.UUID)error{
	insurancer := new(Insurancer)

	//get old values
	if err := db.First(insurancer,"id = ?",insurancer_id).Error; err != nil{
		log.Println("failed to get insurancer for updating:",err.Error())
		return errors.New("failed to delete insurancer")
	}
	oldValues := insurancer

	user_id,_ := GetAuthUserID(c)
	role := GetAuthUser(c)
	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"Insurancer",insurancer_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}

	return nil
}

//get insurancer by id
func GetInsurancer(c *fiber.Ctx, insurancer_id uuid.UUID)(*Insurancer,error){
	insurancer := new(Insurancer)

	if err := db.First(insurancer,"id = ?",insurancer_id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println("record not found:",err.Error())
			return nil, errors.New("record not found")
		}
		log.Println("error getting insurancer details:",err.Error())
		return nil, errors.New("errors getting insurancer details")
	}

	return insurancer, nil
}