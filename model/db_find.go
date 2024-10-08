package model

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/DANCANKARANI/QVP/database"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
var db =database.ConnectDB()

/*
finds user using phone number only
@params phone_number
*/

func UserExist(c *fiber.Ctx, phoneNumber string) (bool, *User, error) {
    var existingUser User

    // Detailed logging
    log.Printf("Checking for user with phone number: %s", phoneNumber)

    result := db.Where("phone_number = ?", phoneNumber).First(&existingUser)
    if result.Error != nil {
        // Log the detailed error
        log.Printf("Error finding user with phone number %s: %v", phoneNumber, result.Error)

        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return false, nil, nil
        }

        return false, nil, fmt.Errorf("database error: %v", result.Error)
    }
	log.Printf("User found: %+v", existingUser)
    return true, &existingUser, nil
}
/*
updates the reset password code in the database
@params phone_number
@params email
@params reset_code
@paarams expiration_time
*/
func AddResetCode(c *fiber.Ctx,phone_number,email,code string,exp_time time.Time) error {
	user := User{}
	db.AutoMigrate(&user)
	result:=db.Where("phone_number = ? AND email = ?",phone_number,email).First(&user)
	if result.Error != nil {
		return utilities.ShowError(c,"failed to get user",fiber.StatusInternalServerError)
	}
	user.ResetCode=code
	user.CodeExpirationTime=exp_time
	result = db.Save(&user)
	if result.Error != nil {
		return utilities.ShowError(c,"failed to save data",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"code sent",fiber.StatusOK)
}
/*
finds if the user with the given email and phone number is registered
@params email
@params phone_number
*/
func FindUser(email, phoneNumber string) (User, error) {
	user := User{}
	err_str := fmt.Sprintf("user with email %s and phone number %s does not exist", email, phoneNumber)
	err := db.Where("phone_number = ? AND email = ?", phoneNumber, email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			
			return user, errors.New(err_str)
		}
		return user, errors.New(err_str)
	}
	return user, nil
}


/*
finds if the rider with the given email and phone number is registered
@params email
@params phone_number
*/
func FindRider(email, phoneNumber string) (Rider, error) {
	rider := Rider{}
	err_str := fmt.Sprintf("user with email %s and phone number %s does not exist", email, phoneNumber)
	err := db.Where("phone_number = ? AND email = ?", phoneNumber, email).First(&rider).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("")
			return rider, errors.New(err_str)
		}
		return rider, errors.New(err_str)
	}
	return rider, nil
}

/*
finds if the rider exists
@params phone number
*/
func RiderExist(c *fiber.Ctx, phoneNumber string) (bool, *Rider, error) {
   var  existingUser Rider

    // Detailed logging
    log.Printf("Checking for user with phone number: %s", phoneNumber)

    result := db.Where("phone_number = ?", phoneNumber).First(&existingUser)
    if result.Error != nil {
        // Log the detailed error
        log.Printf("Error finding user with phone number %s: %v", phoneNumber, result.Error)

        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("record not found:",result.Error.Error())
            return false, nil, nil
        }

        return false, nil, fmt.Errorf("database error: %v", result.Error)
    }
	log.Printf("User found: %+v", existingUser)
    return true, &existingUser, nil
}


/*
finds dependants using phone number only
@params phone_number
*/
func GetAuthUserID(c *fiber.Ctx)(uuid.UUID,error){
	user_id :=c.Locals("user_id")
	id,ok := user_id.(*uuid.UUID)
	if !ok{
		return uuid.Nil,errors.New("unauthorized")
	}
	user_id=*id
	return user_id.(uuid.UUID),nil
}
func GetAuthUser(c *fiber.Ctx)(string){
	user:= c.Locals("role")
	if user == nil{
		log.Println("empty role")
	}
	role, true := user.(string)
	if !true{
		log.Println("failed to convert",user)
	}
	return role
}

//find user with email
func EmailExist(c *fiber.Ctx, email string) (bool, *User, error) {
    var existingUser User

    // Detailed logging
    log.Printf("Checking for user with email: %s", email)

    result := db.Where("email = ?", email).First(&existingUser)
    if result.Error != nil {
        // Log the detailed error
        log.Printf("Error finding user with email %s: %v", email, result.Error)

        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return false, nil, nil
        }

        return false, nil, fmt.Errorf("database error: %v", result.Error)
    }
	log.Printf("User found: %+v", existingUser)
    return true, &existingUser, nil
}