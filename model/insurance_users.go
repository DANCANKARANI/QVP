package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
var insuranceUser = new(InsuranceUser)

func AssisgnInsuranceUser(c *fiber.Ctx)(*InsuranceUser,error){
	insuranceUser.ID = uuid.New()
	if err := c.BodyParser(&insuranceUser); err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	err:= db.Create(&insuranceUser).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to assing users insurance")
	}
	return insuranceUser, nil 
}

func UpdateInsuranceUser(c *fiber.Ctx, insuranceUserID uuid.UUID) (*InsuranceUser, error) {
    // Parse the request body to get the updated data
    updatedData := new(InsuranceUser)
    if err := c.BodyParser(&updatedData); err != nil {
        log.Println("Error parsing JSON data:", err.Error())
        return nil, errors.New("failed to parse JSON data")
    }

    // Check if the insurance user exists
    var existingInsuranceUser InsuranceUser
    if err := db.First(&existingInsuranceUser, "id = ?", insuranceUserID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println("Insurance user not found:", err.Error())
            return nil, errors.New("insurance user not found")
        }
        log.Println("Error finding insurance user:", err.Error())
        return nil, errors.New("failed to find insurance user")
    }

    // Update the existing record with the new data
    if err := db.Model(&existingInsuranceUser).Updates(updatedData).Error; err != nil {
        log.Println("Error updating insurance user:", err.Error())
        return nil, errors.New("failed to update insurance user")
    }

    return &existingInsuranceUser, nil
}


func DeleteInsuranceUser(insurance_user_id uuid.UUID)error{
	err :=db.First(&insuranceUser,"id = ?", insurance_user_id).Delete(&insuranceUser).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete insurance user")
	}
	return nil
}

func GetUsersWithInsurance()(*ResponseUser,error){
	user := new(User)
	err := db.Preload("Insurance").Find(&user).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to get users with insurance")
	}
	return &ResponseUser{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
	},nil
}

