package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Vat value
var vatRate = 16.0
/*
Adds prescription 
@params c *friber.Ctx
@params user_id
*/
func AddPrescription(c *fiber.Ctx,user_id uuid.UUID) (*Prescription, error) {
	db.AutoMigrate(&Prescription{})
	body := Prescription{}
	if err:=c.BodyParser(&body);err != nil {
		log.Fatal(err.Error())
		return nil, errors.New("failed to add prescription")
	}
	prescription :=&Prescription{
		SubTotal: body.SubTotal,
	}
	prescription.CalculateVAT(vatRate)
	body.SubTotal = prescription.SubTotal
	body.VAT=prescription.VAT
	body.Total=prescription.Total
	body.UserApprovedBy = user_id
	body.UserValidatedBy = user_id
	body.ID = uuid.New()
	err:=db.Create(&body).Error
	if err != nil {
		return nil, errors.New("failed to add data")
	}
	return &body,nil
}
/*
Gets users prescriptions
@params id
*/
func GetUsersPrescription(c *fiber.Ctx, id uuid.UUID) (*[]Prescription, error) {
    response := new([]Prescription)
    err := db.Preload("User").First(&response, "user_approved_by = ?", id).Error
    if err != nil {
        
        log.Println(err.Error())
        return nil, errors.New("failed to get prescriptions")
    }
    return response, nil
}

/*
updates the prescription
@params id
@params user_id
*/
func UpdatePrescription(c *fiber.Ctx,id,user_id uuid.UUID)(*Prescription,error){
	body := Prescription{}
	prescription := Prescription{
		SubTotal:body.SubTotal,
	}

	if err:=c.BodyParser(&body); err != nil{
		return nil,errors.New("failed to parse data")
	}
	prescription.CalculateVAT(vatRate)
	body.SubTotal = prescription.SubTotal
	body.VAT=prescription.VAT
	body.Total=prescription.Total
	body.UserApprovedBy=user_id
	err := db.First(&Prescription{},"id = ?",id).Updates(&body).Scan(&body).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update prescription")
	}
	return &body,nil
}

/*
Deletes the prescription
@params id
*/
func DeletePrescription(c *fiber.Ctx, id uuid.UUID) error {
    prescription := Prescription{}
    
    // Check if the prescription exists
    err := db.First(&prescription, "id = ?", id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Printf("No prescription found with ID %s", id)
            return fiber.NewError(fiber.StatusNotFound, "No prescription found")
        }
        log.Printf("Error finding prescription with ID %s: %v", id, err)
        return errors.New("failed to find prescription")
    }
    
    // Attempt to delete the prescription
    err = db.Delete(&prescription).Error
    if err != nil {
        log.Printf("Error deleting prescription with ID %s: %v", id, err)
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete prescription")
    }
    return nil
}

/*
VAT calculator method
@params vatRate
*/
func (p *Prescription)CalculateVAT(vatRate float64){
	p.VAT = p.SubTotal*vatRate/100
	p.Total = p.SubTotal+p.VAT
}