package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
/*
finds if the dependant already exists using the phone number
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

func AddDependant(c *fiber.Ctx)error{
	id := uuid.New()
	body := Dependant{}
	if err := c.BodyParser(&body); err != nil{
		return errors.New("failed to parse json data")
	}
	//getting the user id using GetAuthUerID fuction
	user_id,err := GetAuthUserID(c)
	if err != nil{
		return err
	}
	body.ID=id
	body.UserID=user_id
	result := db.Create(&body)
	if result.Error != nil{
		return errors.New("failed to add dependant: "+result.Error.Error())
	}
	return nil
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
 func UpdateDependant(c *fiber.Ctx, dependant_id string)(*ResponseDependant,error){
	body := Dependant{}
	if err := c.BodyParser(&body); err != nil {
		return &ResponseDependant{},errors.New("failed to parse json data")
	}
	
	result := db.Model(&Dependant{}).Where("id = ?", dependant_id).Updates(&body)
	if result.Error != nil {
		return &ResponseDependant{},result.Error
	}
	response := ResponseDependant{}
	db.First(&Dependant{}).Where("id = ?",dependant_id).Scan(&response)
	return &response,nil
 }

 /*
 deletes the dependant
 @params dependant_id
 @params c *fiber.ctx
 */
func DeleteDependant(c *fiber.Ctx,dependant_id uuid.UUID)error{
	dependant := Dependant{}
	result :=db.First(&dependant,"id = ?",dependant_id)
	if result.Error != nil{
		return errors.New("failed to delete the dependant: "+result.Error.Error())
	}
	db.Delete(&dependant)
	return nil
}

