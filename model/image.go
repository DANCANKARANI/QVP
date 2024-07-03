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
	id :=uuid.New()
	image := Image{
		ID : id,
		OriginalName: file.Filename,
		Type: file.Header.Get("Content-Type"),
	}
	savePath := SanitizeFileName(file.Filename)
	if err := c.SaveFile(file,"./uploads"+savePath); err != nil{
		return nil,errors.New("failed to save the file")
	}
	image.Path = "./uploads/"+file.Filename
	db.AutoMigrate(&Image{})
	if err:= db.Create(&image).Error; err !=nil {
		return nil,errors.New("failed to store the image:"+err.Error())
	}
	return &image,nil
}

func UpdateProfilePhoto(image *Image,id uuid.UUID)error{
	var user User
	//user.ImageID=image.ID
	if err := db.First(&user, "id = ?", id).Updates(&user).Error; err != nil {
		return errors.New("failed to update profile"+err.Error())
	}
	return nil
}