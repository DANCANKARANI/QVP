package admin

import (
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

func RegisterAdminHandler(c *fiber.Ctx)error{
	//parse request body
	body := model.Admin{}
	if err := c.BodyParser(&body); err != nil{
		log.Println("error parsing request body:",err.Error())
		return utilities.ShowError(c,"error parsing request body", fiber.StatusInternalServerError)
	}
	response, err := model.AddAdmin(c,body)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}

	return utilities.ShowSuccess(c,"Admin account created successfully",fiber.StatusOK,response)
}