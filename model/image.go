package model

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"

	firebase "firebase.google.com/go/v4"
	//"firebase.google.com/go/v4/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/option"
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
	image.URL = "./uploads/"+file.Filename
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


func UploadeToFirebaseStorage(filePath,bucketName string)(string, error) {
	ctx := context.Background()
	//initialize firebase app
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err:= firebase.NewApp(ctx,nil,sa)
	if err != nil {
		return "",fmt.Errorf("error initializing the app: %v",err)
	}

	//get storage client
	client,err := app.Storage(ctx)
	if err != nil {
		return "",fmt.Errorf("error getting storage client: %v",err)
	}
	//get a reference to the storage client
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "",fmt.Errorf("error getting bucket: %v",err)
	}

	//OpenFile
	file, err := os.Open(filePath)
	if err != nil {
		return "",fmt.Errorf("error opening file: %v",err)
	}
	defer file.Close()
	uniqueName := uuid.New().String() + filepath.Ext(filePath)

    // Create a reference to the object you want to upload
    object := bucket.Object("images/" + uniqueName)

	//upload the file
	writer :=object.NewWriter(ctx)
	if _,err = io.Copy(writer,file); err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	//get the download URL
	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting object attributes: %v", err)
	}
	downloadURL := attrs.Metadata["access token"]
	return downloadURL,nil
}

func InitFirebase(){
	filePath := "C:/QVP/uploads/2024-07-01 10_51_54.869511 +0300 EAT m=+17.424147001_6.jpg"
	bucketName := "chat-f427d.appspot.com"
	downlaodURL,err := UploadeToFirebaseStorage(filePath,bucketName)
	if err != nil {
		fmt.Println("error:",err)
	}else{
		fmt.Println("Image uploaded to:",downlaodURL)
	}
}