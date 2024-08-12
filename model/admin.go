package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	exist, err := AdminExist(c, body.PhoneNumber)
	if exist && err != nil{
		return nil, err
	}
	response := new(ResAdmin)
	if err := db.Create(&body).Scan(&response).Error; err != nil{
		log.Println("error adding aamin:",err.Error())
		return nil, errors.New("failed to add admin")
	}
	return response, nil
}

func AdminExist(c *fiber.Ctx, phone_number string)(bool,error){
	admin := Admin{}
	result := db.First(admin,"phone_number=?",phone_number)
	if result.Error != nil{
		log.Println("error finding admin existence:",result.Error.Error())
		return true, errors.New("error checking admin existence")
	}
	if result != nil{
		err_str := "admin with phone number:"+phone_number+" already exists"
		return true, errors.New(err_str)
	}
	return false, nil
}