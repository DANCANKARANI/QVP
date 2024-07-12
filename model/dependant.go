package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
type ResponseUpdateDependant struct {
	ID		uuid.UUID			`json:"id"`
	FullName 	string			`json:"full_name"`
	PhoneNumber string 			`json:"phone_number"`
	Relationship string 		`json:"relationship"`
	MemberNumber string 		`json:"member_number"`
	Status 		string 			`json:"status"`
	InsuranceID uuid.UUID		`json:"insurance_id"`
	UserID	uuid.UUID			`json:"user_id"`
}

func DependantExist(c *fiber.Ctx, member_number string) (bool, *Dependant, error) {
	existingDependant := Dependant{}
	err := db.Where("member_number = ?", member_number).First(&existingDependant).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Dependant not found
			return false, nil,nil
		}
		// Other errors
		log.Println("Error checking existence of dependant")
		return false, nil, err
	}
	// Dependant found
	err_string:="dependant with member number "+member_number+" exists"
	return true, &existingDependant, errors.New(err_string)
}

//add deoendant
func AddDependant(c *fiber.Ctx)error{
	id := uuid.New()
	body := Dependant{}
	if err := c.BodyParser(&body); err != nil{
		return errors.New("failed to parse json data")
	}
	//getting the user id using GetAuthUerID fuction
	user_id,err := GetAuthUserID(c)
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to get user id")
	}
	body.ID=id
	body.UserID=user_id
	err = db.Create(&body).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to add dependant")
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
			log.Fatal(err.Error())
        return nil,errors.New("failed to get data")
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
			ImageID:dependant.Insurance.ImageID,
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
 func UpdateDependant(c *fiber.Ctx, dependant_id string)(*ResponseUpdateDependant,error){
	body := Dependant{}
	if err := c.BodyParser(&body); err != nil {
		log.Fatal(err.Error())
		return nil,errors.New("failed to parse json data")
	}
	//update dependant
	err:= db.Model(&Dependant{}).Where("id = ?", dependant_id).Updates(&body).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil,errors.New("failed to update dependant")
	}
	//query response data
	response,err:=GetDependantResponse(dependant_id)
	if err != nil{
		log.Fatal(err.Error())
		return nil,errors.New("failed to get response")
	}
	resDependant:=ResponseUpdateDependant{
		ID:response.ID,
		FullName: response.FullName,
		Relationship: response.Relationship,
		MemberNumber: response.MemberNumber,
		Status: response.Status,
		InsuranceID: response.InsuranceID,
		UserID: response.UserID,
	}
	return &resDependant,nil
 }
//get response to dependants
 func GetDependantResponse(dependantID string) (*Dependant, error) {
    response := Dependant{}
    
    // Query the database to fetch the dependant and preload associated data
    err := db.Preload("User").Preload("Insurance").First(&response, "id = ?", dependantID).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Printf("Dependant with ID %s not found", dependantID)
            return nil, errors.New("dependant not found")
        }
        log.Printf("Error fetching dependant with ID %s: %v", dependantID, err)
        return nil, fmt.Errorf("database error")
    }
    return &response, nil
}

 /*
 deletes the dependant
 @params dependant_id
 @params c *fiber.ctx
 */
 func DeleteDependant(c *fiber.Ctx, dependant_id uuid.UUID) error {
	dependant := Dependant{}
	err := db.First(&dependant, "id = ?", dependant_id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("dependant not found")
		}
		log.Println(err.Error())
		return errors.New("failed to delete dependant")
	}
	if err := db.Delete(&dependant).Error; err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete the dependant")
	}
	return nil
}

