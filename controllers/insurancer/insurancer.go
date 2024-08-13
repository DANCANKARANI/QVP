package insurancer

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

//update insurancer handler
func UpdateInsurancer(c *fiber.Ctx)error{
	insurancer_id, _ := model.GetAuthUserID(c)
	response,err := model.UpdateInsurancer(c,insurancer_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated insurancer details",fiber.StatusOK,response)
}

//get insurancer handler
func GetInsurancerHandler(c *fiber.Ctx)error{
	insurancer_id, _ := model.GetAuthUserID(c)
	response, err := model.GetInsurancer(c,insurancer_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved insurancer details",fiber.StatusOK, response)
}

//delete insurancer handler
func DeleteInsurancerHandler(c *fiber.Ctx)error{
	insurancer_id, _ := model.GetAuthUserID(c)

	if err := model.DeleteInsurancer(c,insurancer_id); err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}

	return utilities.ShowMessage(c,"successfully deleted insurancer",fiber.StatusOK)
}

//update insurancer profile image handler
func UpdateInsurancerProfilePicHandler(c *fiber.Ctx)error{
	insurancer_id, _:= model.GetAuthUserID(c)
	response, err := model.UpdateInsurancerProfilePic(c,insurancer_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"insurancer profile image updated successfully",fiber.StatusOK,response)
}