package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//Vat value
var vatRate = 16.0
/*
Adds prescription 
@params c *friber.Ctx
@params rider_id
@params user_id
@params admin_id
*/
func AddPrescription(c *fiber.Ctx,user_id, rider_id,admin_id uuid.UUID) (*Prescription, error) {
	db.AutoMigrate(&Prescription{})
	body := Prescription{}
	if err:=c.BodyParser(&body);err != nil {
		return nil, errors.New("failed to get data")
	}
	prescription :=&Prescription{
		SubTotal: body.SubTotal,
	}
	prescription.CalculateVAT(vatRate)
	body.SubTotal = prescription.SubTotal
	body.VAT=prescription.VAT
	body.Total=prescription.Total
	body.AdminApprovedBy=admin_id
	body.DeliveredBy = rider_id
	body.UserApprovedBy = user_id
	body.UserValidatedBy = user_id
	body.AdminApprovedBy = admin_id
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
func GetUsersPrescription(c *fiber.Ctx,id string)(*Prescription,error){
	response := Prescription{}
	err := db.Preload("User").Preload("Image").First(&response, "user_approved_by = ?", id).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to get prescriptions")
	}
	return &response,nil
}
/*
updates the prescription
@params id
@params user_id
@params rider_id
@params admin_id
*/
func UpdatePrescription(c *fiber.Ctx,id,user_id,rider_id,admin_id uuid.UUID)(*Prescription,error){
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
	body.DeliveredBy = rider_id
	body.UserValidatedBy=admin_id
	body.AdminApprovedBy = admin_id
	body.AdminValidateBy = admin_id
	err := db.First(&Prescription{},"id = ?",id).Updates(&body).Error
	if err != nil {
		return nil,errors.New("failed to update prescription")
	}
	return &body,nil
}

/*
Deletes the prescription
*/
func DeletePrescription(c *fiber.Ctx, id string) error {
	prescription := Prescription{}
	err :=db.First(&Prescription{}, "id = ?",id).Delete(&prescription).Error
	if err != nil{
		return errors.New("failed to delete prescription")
	}
	return nil
}


func (p *Prescription)CalculateVAT(vatRate float64){
	p.VAT = p.SubTotal*vatRate/100
	p.Total = p.SubTotal+p.VAT
}