package rider

import (
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

//rider signup handler
func CreateRiderAccount(c *fiber.Ctx)error{
	rider := new(model.Rider)

	//parse request body
	if err := c.BodyParser(&rider); err != nil{
		log.Println("error parsing request body:",err.Error())
		return utilities.ShowError(c,"error parsing request body",fiber.StatusInternalServerError)
	}

	//validate email and phone
	valid,err:=model.IsValidData(rider.Email,rider.PhoneNumber)
	if ! valid{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}

	//check if rider exists
	exist,_,err:=model.RiderExist(c,rider.PhoneNumber)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	if exist{
		err_str := "user with phone number "+rider.PhoneNumber+" already exist"
		return utilities.ShowError(c,err_str,fiber.StatusInternalServerError)
	}

	//hash password
	hashed_password,err:=utilities.HashPassword(rider.Password)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	rider.Password = hashed_password

	//create account
	id,err := model.CreateRiderAccount(c,*rider)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	newValues := rider

	//update audit logs
	if err := utilities.LogAudit("Delete",id,"rider","Rider",id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}

	//return response
	return utilities.ShowMessage(c,"rider account created successfully",fiber.StatusOK)
}