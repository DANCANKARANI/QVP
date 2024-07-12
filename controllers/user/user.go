package user

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	response,err := model.GetAllUsersDetails(c)
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
//add user insurance
func AddUserInsurance(c *fiber.Ctx)error{
	user_id,_:=model.GetAuthUserID(c)
	insurance_id,_:=uuid.Parse(c.Params("id"))
	insuranceUser,err:=model.AddUserInsurance(user_id,insurance_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"succssfully added users insurance",fiber.StatusOK,insuranceUser)
}
//updates user insurance
func UpdateUserInsurance(c *fiber.Ctx)error{
	user_id,err:=model.GetAuthUserID(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusUnauthorized)
	}
	id,_:=uuid.Parse(c.Params("id"))
	insurance_id,_:=uuid.Parse(c.Query("insurance_id"))
	insuranceUser,err:=model.UpdateUserInsurance(id,user_id,insurance_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"succssfully updated users insurance",fiber.StatusOK,insuranceUser)
}
