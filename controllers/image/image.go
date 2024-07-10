package image

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
Adds and updates the image
*/
func UploadImagesHandler(c *fiber.Ctx)error{
	file, err := c.FormFile("image")
	if err != nil {
		return utilities.ShowError(c,"failed to upload the image:", fiber.StatusInternalServerError)
	}
	image,err := model.UploadImage(c,file)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.ErrInternalServerError.Code)
	}
	return utilities.ShowSuccess(c,"successfuly uploaded image",fiber.StatusOK,image)
}
/*func UpdateImageHandler(c *fiber.Ctx)error{
	id,_:=uuid.Parse(c.Params("id"))
	image := &model.Image{}
	if err := model.UpdateProfilePhoto(image,id);err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return nil
}*/

//delete the image
func DeleteImageHandler(c *fiber.Ctx)error{
	id,_:= uuid.Parse(c.Params("id"))
	err :=model.DeleteProfilePhoto(id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully removed the profile image",fiber.StatusOK)
}