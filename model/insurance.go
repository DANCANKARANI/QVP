package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
type ResInsurance struct{
	ID 				uuid.UUID 	`json:"id"`
	InsuranceName 	string 		`json:"insurance_name"`
	ImageID 		uuid.UUID 	`json:"image_id"`
	Description 	string 		`json:"description"`			
}
//adds insurance
func AddInsurace(c *fiber.Ctx)error{
	user_id,_:=GetAuthUserID(c)

	role := GetAuthUser(c)

	ID := uuid.New()

	body := Insurance{}
	//get request body
	if err := c.BodyParser(&body); err != nil{
		log.Println("failed to parse request insurance body",err.Error())
		return errors.New("failed to parse request body")
	}

	body.ID=ID

	//create insurance
	err := db.Create(&body).Error
	if err != nil{
		return errors.New("failed to add insurance:"+err.Error())
	}
	newValues := body

	 //update audit logs
	 if err := utilities.LogAudit("Create",user_id,role,"Insurance",ID,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
    
	return nil
}
//updates insurance
func UpdateInsurance(c *fiber.Ctx,insurance_id uuid.UUID)(*Insurance,error){
	user_id, _:= GetAuthUserID(c)

	role := GetAuthUser(c)

	body := Insurance{}

	insurance := new(Insurance)
	
	//get request body
	if err := c.BodyParser(&body); err != nil{
		return nil,errors.New("failed to parse json data:"+err.Error())
	}

	//find insurance record
	if err := db.First(&insurance,"id =?",insurance_id).Error; err !=nil{
		log.Println("error getting insurance record for update", err.Error())
		return nil, errors.New("failed to update insurance")
	}

	oldValues := insurance

	//update insurance model
	if err:=db.Model(&insurance).Updates(&body).Error; err != nil{
		log.Println("error updating insurance",err.Error())
		return nil, errors.New("failed to update insurance")
	}

	newValues := insurance

	 //update audit logs
	 if err := utilities.LogAudit("Update",user_id,role,"Insurance",insurance_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return newValues,nil
}
/*
gets insurance by id
*/
func GetOneInsurance(c *fiber.Ctx)(*Insurance,error){
	id,err:=uuid.Parse(c.Params("id"))
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to get insurance")
	}
	response :=Insurance{}
	err = db.First(&Insurance{},"id = ?",id).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to get insurance")
	}
	return &response,nil
}

/*
gets all insurances
*/
func GetAllInsurances(c *fiber.Ctx)(*[]Insurance,error){
	response := []Insurance{}
	err:=db.Model(&Insurance{}).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to get insurance:"+err.Error())
	}
	
	return &response,nil
}

func DeleteInsurance(c *fiber.Ctx,insurance_id uuid.UUID)error{
	user_id, _ := GetAuthUserID(c)

	role := GetAuthUser(c)

	insurance := new(Insurance)

	if err := db.First(&insurance,"id = ?",insurance_id).Error; err != nil{
		log.Println("error finding insurance for delition",err.Error())
		return errors. New("failed to delete insurance")
	}
	oldValues := insurance

	err := db.Delete(&insurance).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("error deleting insurance")
	}

	 //update audit logs
	 if err := utilities.LogAudit("Delete",user_id,role,"Insurance",insurance_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
	
	return nil
}