package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
) 
type PaginatedPrescriptionDetails struct {
	Page 				int		`json:"page"`
	PageSize 				int	`json:"page_size"`
	TotalPrescriptions 		int64	`json:"total_prescriptions"`
	TotalPages 				int64	`json:"total_pages"`
	PrscriprionDetail		*[]PrescriptionDetail	`json:"prescription_detail"`
	ResponsePrescriptions 	*[]ResponsePrescription `json:"prescriptions"`
}

// AddPrescriptionDetail adds a new prescription detail
func AddPrescriptionDetail(c *fiber.Ctx, prescriptionID uuid.UUID) (*PrescriptionDetail, error) {
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	prescriptionDetail := new(PrescriptionDetail)

	prescriptionDetail.PrescriptionPath, _ = utilities.GenerateUrl(c, "prescription")

	prescriptionDetail.ClaimPath, _ = utilities.GenerateUrl(c, "claim")

	prescriptionDetail.OtherFormPath, _ = utilities.GenerateUrl(c, "other_form")

	prescriptionDetail.PrescriptionID = prescriptionID
	prescriptionDetail.ID = uuid.New()

	if err := db.Create(prescriptionDetail).Error; err != nil {
		log.Println("Error adding prescription detail:", err)
		return nil, errors.New("failed to add prescription detail")
	}

	newValues := prescriptionDetail

	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"Prescription Detail",prescriptionDetail.ID,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}

	//response
	return prescriptionDetail, nil
}

// UpdatePrescriptionDetails updates prescription details
func UpdatePrescriptionDetails(c *fiber.Ctx, prescriptionDetailID, prescriptionID uuid.UUID) (*PrescriptionDetail, error) {
	user_id, _ := GetAuthUserID(c)

	role := GetAuthUser(c)

	prescriptionDetail := new(PrescriptionDetail)

	
	prescriptionDetail.PrescriptionPath, _= utilities.GenerateUrl(c, "prescription")
	
	prescriptionDetail.ClaimPath, _ = utilities.GenerateUrl(c, "claim")

	prescriptionDetail.OtherFormPath, _ = utilities.GenerateUrl(c, "other_form")
	
	prescriptionDetail.PrescriptionID = prescriptionID

	//find prescription detail
	err := db.First(&prescriptionDetail,"id = ?",prescriptionDetailID).Error
	if err != nil{
		log.Println("error updating prescription details",err.Error())
		return nil, errors.New("failed to update prescription detail")
	}

	oldValues := prescriptionDetail

	// Update the prescription detail
	if err := db.Model(&prescriptionDetail).Updates(prescriptionDetail).Error; err != nil {
		log.Println("Error updating prescription detail:", err)
		return nil, errors.New("failed to update prescription detail")
	}
	
	newValues := prescriptionDetail

	//update audit log
	if err := utilities.LogAudit("Update",user_id,role,"Prescription Detail",prescriptionDetailID,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return prescriptionDetail, nil
}
/*
deletes prescription detail
@params id
*/
func DeletePrescriptionDetail(c *fiber.Ctx,id uuid.UUID)(error){
	user_id, _:= GetAuthUserID(c)

	role := GetAuthUser(c)

	prescriptionDetail :=new(PrescriptionDetail)

	//find record
	err := db.First(prescriptionDetail,"id = ?",id).Error
	if err != nil{
		log.Println("error deleting prescription details", err.Error())
		return errors.New("failed to delete prescription details")
	}
	oldValues := prescriptionDetail

	//delete record
	if err:= db.Delete(&prescriptionDetail).Error; err != nil{
		log.Println(err.Error())
		return errors.New("failed to delete prescription details")
	}

	//update audit log
	if err := utilities.LogAudit("Delete",user_id,role,"Prescription Detail",id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}

	return nil
}
//Get paginated prescription details
func GetUsersPrescriptionDetails(c *fiber.Ctx, userID uuid.UUID) (*[]ResponsePrescription, error) {
	var responsePrescriptions []ResponsePrescription

	// Select specific fields from Prescription and PrescriptionDetail
	err := db.Table("prescriptions").
		Select("prescriptions.id, prescriptions.quote_number, prescriptions.sub_total, prescriptions.vat, prescriptions.total, " +
			"prescription_details.prescription_path, prescription_details.claim_path, prescription_details.other_form_path").
		Joins("left join prescription_details on prescriptions.id = prescription_details.prescription_id").
		Where("prescriptions.user_approved_by = ?", userID).
		Scan(&responsePrescriptions).Error

	if err != nil {
		log.Println("Failed to get users prescription details:", err.Error())
		return nil, errors.New("failed to get users prescription details")
	}

	return &responsePrescriptions, nil
}