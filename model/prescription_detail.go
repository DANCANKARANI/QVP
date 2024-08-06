package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	return prescriptionDetail, nil
}

// UpdatePrescriptionDetails updates prescription details
func UpdatePrescriptionDetails(c *fiber.Ctx, prescriptionDetailID, prescriptionID uuid.UUID) (*PrescriptionDetail, error) {
	prescriptionDetail := new(PrescriptionDetail)

	
	prescriptionDetail.PrescriptionPath, _= utilities.GenerateUrl(c, "prescription")
	
	prescriptionDetail.ClaimPath, _ = utilities.GenerateUrl(c, "claim")

	prescriptionDetail.OtherFormPath, _ = utilities.GenerateUrl(c, "other_form")
	
	prescriptionDetail.PrescriptionID = prescriptionID

	// Update the prescription detail
	if err := db.Model(&PrescriptionDetail{}).Where("id = ?", prescriptionDetailID).
		Updates(prescriptionDetail).Scan(prescriptionDetail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record not found:", err)
			return nil, errors.New("record not found")
		}
		log.Println("Error updating prescription detail:", err)
		return nil, errors.New("failed to update prescription detail")
	}

	return prescriptionDetail, nil
}
/*
deletes prescription detail
@params id
*/
func DeletePrescriptionDetail(id uuid.UUID)(error){
	prescriptionDetail :=new(PrescriptionDetail)
	err := db.First(prescriptionDetail,"id = ?",id).Delete(&prescriptionDetail).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to delete prescription details")
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