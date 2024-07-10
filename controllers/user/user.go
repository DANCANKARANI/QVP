package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
)
//get one user handler
func GetOneUserHandler(c *fiber.Ctx) error {
	user,err := model.GetOneUSer(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"user retrieved successfully",fiber.StatusOK,user)
}

//get all users handler
func GetAllUsersHandler(c *fiber.Ctx)error{
	response,err := model.GetAllUsers(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError) 
	}
	return utilities.ShowSuccess(c,"users retrieved successfully",fiber.StatusOK,response)
}

//update user details handler
func UpdateUserHandler(c *fiber.Ctx)error{
	response,err := model.UpdateUser(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"user updated successfully",fiber.StatusOK,response)
}
