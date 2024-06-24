package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
/*
finds if the dependant already exists using the phone number
*/

type ResponseDependant struct{
	ID		uuid.UUID			`json:"id"`
	FullName 	string			`json:"full_name"`
	PhoneNumber string 			`json:"phone_number"`
	Relationship string 		`json:"relationship"`
	MemberNumber string 		`json:"member_number"`
	Status 		string 			`json:"status"`
	InsuranceID uuid.UUID		`json:"insurance_id"`
	UserID	uuid.UUID			`json:"user_id"`
    User    ResponseUser
    Insurance ResInsurance			
}

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
func GetAllDependants(c *fiber.Ctx,user_id uuid.UUID)(*[]ResponseDependant,error){
	var dependants []Dependant
    if err := db.Preload("User").Preload("Insurance").
        Where("user_id = ?", user_id).Find(&dependants).Error; err != nil {
        return nil,errors.New("failed to get data:"+err.Error())
    }
	var response []ResponseDependant
    for _, dependant := range dependants {
        resUser := ResponseUser{
			ID: dependant.User.ID,
            FullName:    dependant.User.FullName,
            PhoneNumber: dependant.User.PhoneNumber,
            Email:      dependant.User.Email,
        }
		resInsurance :=ResInsurance{
			ID: dependant.Insurance.ID,
			InsuranceName: dependant.Insurance.InsuranceName,
			PhotoPath: dependant.Insurance.PhotoPath,
			Description: dependant.Insurance.Description,
		}
        responseDependant := ResponseDependant{
            ID:           dependant.ID,
            FullName:     dependant.FullName,
            PhoneNumber:  dependant.PhoneNumber,
            Status:       dependant.Status,
            Relationship: dependant.Relationship,
            UserID:       dependant.UserID,
            MemberNumber: dependant.MemberNumber,
            InsuranceID:  dependant.InsuranceID,
            User:         resUser,
			Insurance: resInsurance,
        }
		
        response = append(response, responseDependant)
    }
	return &response,nil
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
		return nil,result.Error
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

