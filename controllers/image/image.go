package image

import (
	"log"
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
func UploadProfileImage(c *fiber.Ctx)error{
	user_id,_:=model.GetAuthUserID(c)
	file,err:=c.FormFile("image")
	if err != nil{
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to upload the image:", fiber.StatusInternalServerError)
	}
	image,err:=model.UploadImage(c,file)
	if err != nil {
		log.Println(err.Error())
	}
	err =model.UpdateUserProfile(user_id,image.ID)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully uploaded the image",fiber.StatusOK)
}
func UploadInsuranceIcon(c *fiber.Ctx)error{
	insurance_id,_:=uuid.Parse(c.Params("id"))
	file,err:=c.FormFile("image")
	if err != nil{
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to upload the image:", fiber.StatusInternalServerError)
	}
	image,err:=model.UploadImage(c,file)
	if err != nil {
		log.Println(err.Error())
	}
	err =model.UpdateInsuranceIcon(insurance_id,image.ID)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully uploaded the insurance icon",fiber.StatusOK)
}

func UploadPaymentMethod(c *fiber.Ctx)error{
	payment_method_id,_:=uuid.Parse(c.Params("id"))
	file,err:=c.FormFile("image")
	if err != nil{
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to upload the image:", fiber.StatusInternalServerError)
	}
	image,err:=model.UploadImage(c,file)
	if err != nil {
		log.Println(err.Error())
	}
	err =model.UpdatePaymentMethodIcon(payment_method_id,image.ID)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully uploaded the insurance icon",fiber.StatusOK)
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