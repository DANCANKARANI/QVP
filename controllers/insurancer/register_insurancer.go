package insurancer

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

//create insurancer account handler
func CreateInsurancerAccountHandler(c *fiber.Ctx) error {
	insurancer := new(model.Insurancer)

	//parse request body
	if err := c.BodyParser(&insurancer); err != nil{
		log.Println("failed  to parse request data:",err.Error())
		return errors.New("failed to parse request data")
	}

	//validate email and phone number
	valid,err :=model.IsValidData(insurancer.Email,insurancer.PhoneNumber)
	if !valid{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	
	//find if insurancer exist
	exist,_,err := model.InsurerExist(c,insurancer.PhoneNumber)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	if exist{
		err_str := "insurancer with phone number "+insurancer.PhoneNumber+" already esxist"
		return utilities.ShowError(c,err_str,fiber.StatusInternalServerError)
	}

	//create account
	id,err := model.CreateInsurancerAccount(c,*insurancer)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInsufficientStorage)
	}

	//update audit logs
	if err := utilities.LogAudit("Register",id,"insurancer","Insurancer",id,nil,insurancer,c); err != nil{
		log.Println(err.Error())
	}

	//response
	return utilities.ShowMessage(c,"insurancer account registered successfully",fiber.StatusOK)
}