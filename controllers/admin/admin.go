package admin

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

//update admin handler
func UpdateAdminHandler(c *fiber.Ctx) error {

	//update admin details
	response, err := model.UpdateAdminDetails(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated admin details",fiber.StatusOK,response)
}

//get admin details
func GetAdminDetailsHandler(c *fiber.Ctx)error{
	admin_id, _ := model.GetAuthUserID(c)

	response, err := model.GetAdminDetails(c,admin_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"admin details retrieved successfully",fiber.StatusOK,response)
}

//get all admins handler
func GetAllAdminsHandler(c *fiber.Ctx)error{
	response, err := model.GetAllAdmins(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved all admins",fiber.StatusOK, response)
}

//delete admin by phone number
func DeleteAdmin(c *fiber.Ctx)error{
	phone_number := c.Query("phone")
	err := model.DeleteAdmnin(c,phone_number)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"admin deleted successfully",fiber.StatusOK)
}

//update profile pic handler
func UpdateProfilePicHandler(c *fiber.Ctx)error{
	admin_id, _:= model.GetAuthUserID(c)
	response,err := model.UpdateProfilePic(c,admin_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated admin profile image",fiber.StatusOK,response)
}