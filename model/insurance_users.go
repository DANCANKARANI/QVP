package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func AssisgnInsuranceUser(c *fiber.Ctx)(*InsuranceUser,error){
	user_id, _:= GetAuthUserID(c)

	role := GetAuthUser(c)

	insuranceUser := new(InsuranceUser)

	insuranceUser.ID = uuid.New()

	//parse request body
	if err := c.BodyParser(&insuranceUser); err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}

	//create insuranceUser
	err:= db.Create(&insuranceUser).Error
	if err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to assing users insurance")
	}
	newValues := insuranceUser

	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"Insurance User",insuranceUser.ID,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}

	//return response
	return insuranceUser, nil 
}

func UpdateInsuranceUser(c *fiber.Ctx, insuranceUserID uuid.UUID) (*InsuranceUser, error) {
	user_id, _ := GetAuthUserID(c)

	role := GetAuthUser(c)

    updatedData := new(InsuranceUser)

	// Parse the request body to get the updated data
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
	oldValues := existingInsuranceUser

    // Update the existing record with the new data
    if err := db.Model(&existingInsuranceUser).Updates(updatedData).Error; err != nil {
        log.Println("Error updating insurance user:", err.Error())
        return nil, errors.New("failed to update insurance user")
    }
	newValues := existingInsuranceUser

	//update audit logs
	if err := utilities.LogAudit("Update",user_id,role,"Insurance User",insuranceUserID,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

    return &existingInsuranceUser, nil
}


func DeleteInsuranceUser(c *fiber.Ctx,insurance_user_id uuid.UUID)error{
	user_id, _ := GetAuthUserID(c)

	role:=GetAuthUser(c)

	insuranceUser := new(InsuranceUser)
	//getting insurance user
	if err :=db.First(&insuranceUser,"id = ?", insurance_user_id).Error; err != nil{
		log.Println("error getting insurance user for deleting",err.Error())
		return errors.New("failed to delete insurance user")
	}
	oldValues := insuranceUser
	//deleting insurance user
	err :=db.Delete(&insuranceUser).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete insurance user")
	}

	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"Insurance User",insurance_user_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
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

