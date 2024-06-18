package model

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/database"
	"main.go/utilities"
)
var db =database.ConnectDB()
/*
finds user using phone number only
@params phone_number
*/
func UserExist(c *fiber.Ctx,phone_number string)(bool,User,error){
	 existingUser := User{}
	result := db.Where("phone_number = ?",phone_number).First(&existingUser)
	if result.Error != nil {
		//user not found
		return false,existingUser,result.Error
	}
	return true,existingUser, nil
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
func FindUser(email, phone_number string)(User,error){
	user := User{}
	result:=db.Where("phone_number = ? AND email = ?",phone_number,email).First(&user)
	if result.Error != nil {
		return user,result.Error
	}
	return user,nil
}

/*
finds dependants using phone number only
@params phone_number
*/
func DependantExist(c *fiber.Ctx,phone_number string)(bool,Dependant,error){
	existingDependant := Dependant{}
   result := db.Where("phone_number = ?",phone_number).First(&existingDependant)
   if result.Error != nil {
	   //user not found
	   return false,existingDependant,result.Error
   }
   return true,existingDependant, nil
}

/*
get all the dependants for a specific user
*/
func GetAllDependants(c *fiber.Ctx,user_id uuid.UUID)([]Dependant,error){
	existingDependants := []Dependant{}
	result := db.Preload("User").Where("user_id = ?",user_id).Find(&existingDependants)
	if result.Error !=nil{
		return existingDependants,result.Error
	}
	return existingDependants,nil
}


/*
update the dependant details
@params c *fiber.Ctx
*/
func GetDependantID(c *fiber.Ctx)(uuid.UUID,error){
	dependant :=Dependant{}
	if err := c.BodyParser(&dependant);err !=nil {
		return  uuid.Nil,err
	}
	result :=db.Where("member_number =?",dependant.MemberNumber).First(&dependant)
	if result.Error != nil{
		return uuid.Nil,result.Error
	}
	
	return dependant.ID,nil
 }

 /*
 updates the dependants details
 @params dependant_id
 */
 func UpdateDependant(c *fiber.Ctx, dependant_id string)(*Dependant,error){
	body := Dependant{}
	if err := c.BodyParser(&body); err != nil {
		return &Dependant{},errors.New("failed to parse json data")
	}
	
	result := db.Model(&Dependant{}).Where("id = ?", dependant_id).Updates(&body)
	if result.Error != nil {
		return &Dependant{},result.Error
	}
	response := ResponseDependant{}
	db.First(&Dependant{}).Where("id = ?",dependant_id).Scan(&response)
	return &response,nil
 }

 