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
type ResInsurancer struct{
	ID			uuid.UUID 	`json:"id"`
	FullName	string		`json:"full_name"`
	Email		string		`json:"email"`
	PhoneNumber	string		`json:"phone_number"`
	ProfilePhotoPath string	`json:"profile_photo_path"`
}
/*
creates insurancer account
@params body
*/
func CreateInsurancerAccount(c *fiber.Ctx,body Insurancer)(uuid.UUID,error){
	Insurancer :=body

	//hash password
	hashed_password,err := utilities.HashPassword(Insurancer.Password)
	if err != nil{
		return uuid.Nil,errors.New(err.Error())
	}

	Insurancer.Password = hashed_password
	Insurancer.ID=uuid.New()
	//create insurancer
	if err := db.Create(&Insurancer).Error; err != nil{
		log.Println("error creating insurancer account:",err.Error())
		return uuid.Nil,errors.New("failed to create insurancer account")
	}

	return Insurancer.ID,nil
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

func UpdateInsurancer(c *fiber.Ctx, insurancer_id uuid.UUID)(*ResInsurancer,error){
	insurancer := new(Insurancer)
	body := Insurancer{}

	//parse request
	if err :=  c.BodyParser(&body); err != nil{
		log.Println("error parsing request body:", err.Error())
		return nil, errors.New("error parsing request body")
	}
	//validate phone number
	if body.PhoneNumber !=""{
		_,err :=utilities.ValidatePhoneNumber(body.PhoneNumber,country_code)
		if err != nil{
			return nil, err
		}
		exist,_,err:=InsurerExist(c,body.PhoneNumber)
		if err != nil{
			return nil, err
		}
		if exist{
			err_str := "user with phone:"+body.PhoneNumber+" already exist"
			return nil, errors.New(err_str)
		}
	}

	//validate email
	if body.Email !=""{
		_, err := utilities.ValidateEmail(body.Email)
		if err != nil{
			return nil, err
		}
	}

	//hash password
	if body.Password != ""{
		hashed_password, err:= utilities.HashPassword(body.Password)
		if err != nil{
			return nil,err
		}
		body.Password= hashed_password
	}

	//find old values
	if err := db.First(insurancer,"id = ?",insurancer_id).Error; err != nil{
		log.Println("error getting insurancer for update:",err.Error())
		return nil, errors.New("failed to update insurancer")
	}
	oldValues := insurancer

	response := new(ResInsurancer)
	//update insurancer
	if err := db.Model(&insurancer).Updates(body).Scan(&response).Error; err != nil{
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

	return response, nil
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
func GetInsurancer(c *fiber.Ctx, insurancer_id uuid.UUID)(*ResInsurancer,error){
	insurancer := new(Insurancer)
	response := new(ResInsurancer)

	if err := db.First(insurancer,"id = ?",insurancer_id).Scan(&response).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println("record not found:",err.Error())
			return nil, errors.New("record not found")
		}
		log.Println("error getting insurancer details:",err.Error())
		return nil, errors.New("errors getting insurancer details")
	}

	return response, nil
}

/*
updates insurancer profile image
@params insurancer_id
*/
func UpdateInsurancerProfilePic(c *fiber.Ctx, insurancer_id uuid.UUID)(*ResInsurancer,error){
	insurancer := new(Insurancer)

	//generate image url
	profile_photo_path,err:=utilities.GenerateUrl(c,"profile")
	if err != nil{
		return nil, err
	}
	Insurancer := Insurancer{
		ProfilePhotoPath: profile_photo_path,
	}
	
	//find user
	if err := db.First(&insurancer,"id = ?",insurancer_id).Error; err != nil{
		log.Println("insurancer not found:",err.Error())
		return nil, errors.New("failed to update profile image")
	}
	oldValues := insurancer
	response := new(ResInsurancer)
	//update profile image
	if err := db.Model(insurancer).Updates(&Insurancer).Scan(response).Error; err != nil{
		log.Println("failed to update profile image:", err.Error())
		return nil, errors.New("failed to update profile image")
	}
	newValues := insurancer
	
	//update audit log

	role := GetAuthUser(c)

	if err := utilities.LogAudit("Update",insurancer_id,role,"Insurancer",insurancer_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return response, nil
}