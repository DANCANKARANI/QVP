package prescription

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//add prescription handler
func AddPrescriptionHandler(c *fiber.Ctx)error{
	user_id,_:= model.GetAuthUserID(c)
	//user_id,_ :=uuid.Parse(c.Query("user_id"))
	//rider_id,_ :=uuid.Parse(c.Query("rider_id"))
	prescription,err := model.AddPrescription(c,user_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"Prescription added successfully",fiber.StatusOK,prescription)
}
//get prescrption handler
func GetPrescriptionsHandler(c *fiber.Ctx)error{
	id,_ := model.GetAuthUserID(c)
	//id := c.Query("user_id")
	response,err:=model.GetUsersPrescription(c,id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println(err.Error())
            return utilities.ShowError(c,"prescriptions not found",fiber.StatusNotFound)
        }
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"prescription retrieved successfully",fiber.StatusOK,response)
}
//update prescription handler
func UpdatePrescriptionHandler(c *fiber.Ctx)error{
	id,_:= uuid.Parse(c.Params("id"))
	user_id,_:= model.GetAuthUserID(c)
	response,err := model.UpdatePrescription(c,id,user_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return utilities.ShowMessage(c,"prescription not found",fiber.StatusNotFound)
		}
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated prescription",fiber.StatusOK,response)
}
//Delete prescription handler
func DeletePrescriptionHandler(c *fiber.Ctx)error{
	id,_:= uuid.Parse(c.Params("id"))
	err := model.DeletePrescription(c,id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return utilities.ShowMessage(c,"prescription not found",fiber.StatusNotFound)
		}
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"prescription deleted successfully",fiber.StatusOK)
}

//gets all the pagineted prescriptions
func GetAllPrescriptionsHandler(c *fiber.Ctx)error{
	response,code,err:=model.GetPaginatePrescriptions(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),code)
	}
	return utilities.ShowSuccess(c,"successfully retrieved presicriptions",code,response)
}
//update prescription detail handler
func UpdatePrescriptionDetailHandler(c *fiber.Ctx)error{
	prescription_detail_id, _:=uuid.Parse(c.Params("id"))
	log.Println(prescription_detail_id)
	prescription_id, _:=uuid.Parse(c.Params("prescription_id"))
	response, err := model.UpdatePrescriptionDetails(c,prescription_detail_id,prescription_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated user's prescription details",fiber.StatusOK,response)
}
//add prescription details handler
func AddPrescriptionDetailHandler(c *fiber.Ctx)error{
	prescription_id,_ := uuid.Parse(c.Params("prescription_id"))
	log.Println(prescription_id)
	response, err := model.AddPrescriptionDetail(c,prescription_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added prescription details",fiber.StatusOK,response)
	 
}
//gets users prescription details
func GetUsersPrescriptionDetailHandler(c *fiber.Ctx)error{
	user_id,_:= model.GetAuthUserID(c)
	response, err := model.GetUsersPrescriptionDetails(c,user_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved users prescription details",fiber.StatusOK,response)
}

