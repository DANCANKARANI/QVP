package model

import (
	"errors"
	"mime/multipart"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func SanitizeFileName(fileName string) string {
    // Define a regular expression to match invalid characters
    invalidChars := regexp.MustCompile(`[<>:"/\\|?*]+`)
    return invalidChars.ReplaceAllString(fileName, "_")
}

func UploadImage(c *fiber.Ctx,file *multipart.FileHeader)(*Image,error){

	image := Image{
		OriginalName: file.Filename,
		Type: file.Header.Get("Content-Type"),
	}
	savePath := SanitizeFileName(file.Filename)
	if err := c.SaveFile(file,"./uploads"+savePath); err != nil{
		return nil,errors.New("failed to save the file")
	}
	image.Path = "./uploads/"+file.Filename
	if err:= db.Create(&Image{}); err !=nil {
		return nil,errors.New("failed to store the image")
	}
	return &image,nil
}

func UpdateProfilePhoto(image *Image,id uuid.UUID)error{
	var user User
	user.ImageID=image.ID
	if err := db.First(&user, "id = ?", id).Updates(&user); err != nil {
		return errors.New("failed to update profile")
	}
	return nil
}