package image
import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/DANCANKARANI/QVP/model"
)
func UploadImages(c *fiber.Ctx)error{
	id,_:=model.GetAuthUserID(c)
	file, err := c.FormFile("image")
	if err != nil {
		return utilities.ShowError(c,"failed to upload the image:", fiber.StatusInternalServerError)
	}
	image,err := model.UploadImage(c,file)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.ErrInternalServerError.Code)
	}
	if err := model.UpdateProfilePhoto(image,id);err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfuly uploaded image",fiber.StatusOK,image)
}