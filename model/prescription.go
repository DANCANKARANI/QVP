package model

import (
	"errors"
	"log"
	"strconv"

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
type ResponsePrescription struct {
	ID 			uuid.UUID   `json:"id"`
	QuoteNumber string 		`json:"quote_number"`
	SubTotal  	float64		`json:"sub_total"`
	VAT			float64		`json:"vat"`
	Total 		float64		`json:"total"`
	PrescriptionPath string 	`json:"prescription_path"`
	ClaimPath	string		`json:"claim_path"`
	OtherFormPath string	`json:"other_form_path"`
}
func AddPrescription(c *fiber.Ctx,user_id uuid.UUID) (*ResponsePrescription, error) {
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
	responsePrescrption:=&ResponsePrescription{
		ID: body.ID,
		QuoteNumber: body.QuoteNumber,
		SubTotal: body.SubTotal,
		VAT: body.VAT,
		Total: body.Total,
	}
	return responsePrescrption,nil
}
/*
Gets users prescriptions
@params id
*/
func GetUsersPrescription(c *fiber.Ctx, id uuid.UUID) (*[]ResponsePrescription, error) {
    response := new([]ResponsePrescription)
    err := db.Preload("User").First(&Prescription{}, "user_approved_by = ?", id).Scan(&response).Error
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
func UpdatePrescription(c *fiber.Ctx,id,user_id uuid.UUID)(*ResponsePrescription,error){
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
	response := new(ResponsePrescription)
	err := db.First(&Prescription{},"id = ?",id).Updates(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update prescription")
	}
	return response,nil
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
gets paginated prescriptions
*/
type PaginatedPrescriptions struct {
	Page int		`json:"page"`
	PageSize int	`json:"page_size"`
	TotalPrescriptions int64	`json:"total_prescriptions"`
	TotalPages int64	`json:"total_pages"`
	ResponsePrescriptions *[]ResponsePrescription `json:"prescriptions"`
}
func GetPaginatePrescriptions(c *fiber.Ctx)(*PaginatedPrescriptions,int,error){
	//get page
	page, err := strconv.Atoi(c.Query("page"))
	if err !=nil || page < 1 {
		log.Println(err.Error())
		return nil,fiber.StatusBadRequest,errors.New("invalid page number")
	}
	//get page size
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize < 1 {
		log.Println(err.Error())
		return nil, fiber.StatusBadRequest,errors.New("invalid page size")
	}
	//response user
	responsePrescription := new([]ResponsePrescription)
	prescriptions := new([]Prescription)
	var totalPages int64
	//get prescription and count
	db.Model(&Prescription{}).Count(&totalPages)
	db.Offset((page -1) * pageSize).Limit(pageSize).Find(&prescriptions).Scan(&responsePrescription)

	paginatedPrescriptions := &PaginatedPrescriptions{
		Page: page,
		PageSize: pageSize,
		TotalPrescriptions: totalPages,
		TotalPages:  (totalPages + int64(pageSize) - 1) / int64(pageSize),
		ResponsePrescriptions: responsePrescription,
	}
	return paginatedPrescriptions,fiber.StatusOK,nil
}

/*
VAT calculator method
@params vatRate
*/
func (p *Prescription)CalculateVAT(vatRate float64){
	p.VAT = p.SubTotal*vatRate/100
	p.Total = p.SubTotal+p.VAT
}