package model

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
type ResInsurance struct{
	ID uuid.UUID `json:"id"`
	InsuranceName string `json:"insurance_name"`
	PhotoPath string 	`json:"photo_path"`
	Description string `json:"description"`			
}
//adds insurance
func AddInsurace(c *fiber.Ctx)error{
	ID := uuid.New()
	body := Insurance{}
	if err := c.BodyParser(&body); err != nil{
		return errors.New("failed to parse the Json data")
	}
	body.ID=ID
	err := db.Create(&body).Error
	if err != nil{
		return errors.New("failed to add insurance:"+err.Error())
	}
	return nil
}
//updates insurance
func UpdateInsurance(c *fiber.Ctx)(*Insurance,error){
	body := Insurance{}
	insurance_id,err := uuid.Parse(c.Query("insurance_id"))
	if err != nil {
		return nil, errors.New("failed to convert id into uuid:"+err.Error())
	}
	if err := c.BodyParser(&body); err != nil{
		return nil,errors.New("failed to parse json data:"+err.Error())
	}
	err = db.Model(&Insurance{}).Where("id = ?",insurance_id).Updates(&body).Error
	if err != nil {
		return nil,errors.New("failed to update insurance:"+err.Error())
	}
	err = db.First(&Insurance{},"id = ?",insurance_id).Scan(&body).Error
	if err != nil {
		return nil,errors.New("failed to get updated row:"+err.Error())
	}
	return &body,nil
}
/*
gets insurance by id
*/
func GetOneInsurance(c *fiber.Ctx)(*Insurance,error){
	id,err:=uuid.Parse(c.Query("id"))
	if err != nil {
		return nil,errors.New("failed to convert id string to uuid:"+err.Error())
	}
	response :=Insurance{}
	err = db.First(&Insurance{},"id = ?",id).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to get insurance:"+err.Error())
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

func DeleteInsurance(c *fiber.Ctx,id string)error{
	insurance := Insurance{}
	err := db.First(&Insurance{},"id = ?",id).Delete(&insurance)
	if err != nil{
		return errors.New("error deleting insurance")
	}
	return nil
}