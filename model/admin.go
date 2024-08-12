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

type ResAdmin struct {
	ID 			uuid.UUID 		`json:"id"`
	FullName	string			`json:"full_name"`
	Email		string			`json:"email"`
	PhoneNumber	string			`json:"phone_number"`
	ProfilePhotoPath	string	`json:"profile_photo_path"`
}

func AddAdmin(c *fiber.Ctx, body Admin)(*ResAdmin,error){

	//valid email and phone
	valid,err:=IsValidData(body.Email,body.PhoneNumber)
	if ! valid && err != nil{
		return nil, errors.New(err.Error())
	}

	//check if admin already exist
	exist,_, err := AdminExist(c, body.PhoneNumber)
	if exist && err != nil{
		return nil, err
	}
	response := new(ResAdmin)
	
	body.ID= uuid.New()
	if body.Password == "" || body.Email==""||body.PhoneNumber =="" ||body.FullName==""{
		return nil, errors.New("fill all credintials")
	}
	hashed_password,_:=utilities.HashPassword(body.Password)
	body.Password = hashed_password

	if err := db.Create(&body).Scan(&response).Error; err != nil{
		log.Println("error adding aamin:",err.Error())
		return nil, errors.New("failed to add admin")
	}

	role := GetAuthUser(c)

	//update log audit
	if err := utilities.LogAudit("Register",body.ID,role,"Admin",body.ID,nil,body,c); err != nil{
		log.Println(err.Error())
	}

	return response, nil
}

/*
checks if the admin exists in the db
@params phone_number
*/
func AdminExist(c *fiber.Ctx, phoneNumber string)(bool,*Admin,error){
	admin := Admin{}
	log.Printf("Checking for user with phone number: %s", phoneNumber)
 
	result := db.Where("phone_number = ?", phoneNumber).First(&admin)
	if result.Error != nil {
		// Log the detailed error
		log.Printf("Error finding user with phone number %s: %v", phoneNumber, result.Error)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("record not found:",result.Error.Error())
			return false, nil, nil
		}

	   return false, nil, fmt.Errorf("database error: %v", result.Error)
   }
   log.Printf("User found: %+v", admin)
   return true, &admin, nil
}

/*\
updates admin details
*/
func UpdateAdminDetails(c *fiber.Ctx)(*ResAdmin,error){
	admin_id, _ := GetAuthUserID(c)
	body := Admin{}

	//parse request body
	if err := c.BodyParser(&body); err != nil{
		log.Println("error parsing request body:", err.Error())
		return nil,errors.New("error parsing request body")
	}

	//validate phone
	if body.PhoneNumber != ""{
		phone_number,err:=utilities.ValidatePhoneNumber(body.PhoneNumber,"KE")
		if err != nil{
			return nil, err
		}
		body.PhoneNumber = phone_number
	}

	//validate email 
	if body.Email != ""{
		email, err:= utilities.ValidateEmail(body.Email)
		if err != nil{
			return nil, err
		}
		body.Email = *email
	}

	//hash password
	if body.Password != ""{
		hashed_password, err := utilities.HashPassword(body.Password)
		if err != nil{
			return nil, err
		}
		body.Password = hashed_password
	}

	//check of user exist
	existingAdmin := new(Admin)
	if err := db.First(&existingAdmin,"id = ?",admin_id).Error; err != nil{
		log.Println("error findind admin:",err.Error())
		return nil, errors.New("failed to update admin")
	}
	oldValues := existingAdmin
	response := new(ResAdmin)
	//update admin model 
	if err := db.Model(existingAdmin).Updates(&body).Scan(response).Error; err != nil{
		log.Println("error updating admin details:",err.Error())
		return nil, errors.New("failed to update admin details")
	}
	newValues := existingAdmin

	role := GetAuthUser(c)

	if err := utilities.LogAudit("Update",admin_id,role,"Admin",admin_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return response, nil
}


//gets admin details by id
func GetAdminDetails(c *fiber.Ctx, admin_id uuid.UUID)(*ResAdmin,error){
	admin := new(Admin)
	response := new(ResAdmin)
	//get admin details
	if err := db.First(admin,"id = ?", admin_id).Scan(response).Error; err != nil{
		log.Println("error getting admin details:", err.Error())
		return nil, errors.New("failed to get admin details")
	}

	return response, nil
}