package model

import (
	"errors"
	"log"
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
	user_id,_:=GetAuthUserID(c)
	id :=uuid.New()
	image := Image{
		ID : id,
		OriginalName: file.Filename,
		Type: file.Header.Get("Content-Type"),
		UserID: user_id,
	}
	savePath := SanitizeFileName(file.Filename)
	if err := c.SaveFile(file,"./uploads"+savePath); err != nil{
		return nil,errors.New("failed to save the file")
	}
	image.Path = "./uploads/"+file.Filename
	db.AutoMigrate(&Image{})
	//add image and update the user
	if err:= db.Create(&image).Error; err !=nil {
		return nil,errors.New("failed to store the image: "+err.Error())
	}
	
	return &image,nil
}



/*
updates user
@params user_id
@params image_id
*/
func UpdateUserProfile(user_id,image_id uuid.UUID)(error){
	body:=User{}
	body.ImageID=image_id
	err := db.First(&User{},"id = ?",user_id).Updates(&body).Scan(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("error in updating the user:"+err.Error())
	}
	return nil
}
/*
updates payment method icon
@params payment_method_id
@params image_id
*/
func UpdatePaymentMethodIcon(payment_method_id,image_id uuid.UUID)error{
	body:=PaymentMethod{}
	body.ImageID=image_id
	err := db.First(&body,"id = ?",payment_method_id).Updates(&body).Scan(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("error in updating the payment method icon")
	}
	return nil
}
/*
updates insurance icon
@params insurance_id
@params image id
*/
func UpdateInsuranceIcon(insurance_id,image_id uuid.UUID)error{
	body:=Insurance{}
	body.ImageID=image_id
	err := db.First(&body,"id = ?",insurance_id).Updates(&body).Scan(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("error in updating the insurance")
	}
	return nil
}


/*
update the profile image
@params image *Image
@params image_id
*/
func UpdateProfilePhoto(image *Image,image_id uuid.UUID)error{
	//user.ImageID=image.ID
	if err := db.First(&image, "id = ?", image_id).Updates(&image).Error; err != nil {
		return errors.New("failed to update profile"+err.Error())
	}
	return nil
}
/*
deletes the profile image
@params image_id
*/
func DeleteProfilePhoto(image_id uuid.UUID)error{
	image := Image{}
	err := db.First(&image, "id =  ?", image_id).Delete(&image).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to remove profile image")
	}
	return nil
}