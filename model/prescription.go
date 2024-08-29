package model

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/DANCANKARANI/QVP/utilities"
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
	role := GetAuthUser(c)
	body := Prescription{}
	
	body.QuoteNumber = GenerateQuoteNumber()

	    // Calculate subtotal and total (for simplicity, this example assumes quote details are added first)
		var quoteDetails []QuoteDetail
		if err := db.Where("prescription_id = ?", body.ID).Find(&quoteDetails).Error; err != nil {
			return nil,errors.New("failed to retrieve quote details")
		}
	
		var subTotal float64
		for _, detail := range quoteDetails {
			subTotal += detail.Total
		}
		body.SubTotal = subTotal

		//calculate vat
	body.CalculateVAT(vatRate)


	body.UserApprovedBy = &user_id
	body.UserValidatedBy = &user_id
	body.ID = uuid.New()
	err:=db.Create(&body).Error
	if err != nil {
		log.Println("error adding prescription:",err.Error())
		return nil, errors.New("failed to add data")
	}
	responsePrescrption:=&ResponsePrescription{
		ID: body.ID,
		QuoteNumber: body.QuoteNumber,
		SubTotal: body.SubTotal,
		VAT: body.VAT,
		Total: body.Total,
	}
	//update audit log
	if err := utilities.LogAudit("Create",user_id,role,"Prescription",body.ID,nil,responsePrescrption,c); err != nil{
		log.Println(err.Error())
	}
	return responsePrescrption,nil
}
/*
Gets users prescriptions
@params id
*/
func GetUsersPrescription(c *fiber.Ctx, id uuid.UUID) (*[]ResponsePrescription, error) {
    response := new([]ResponsePrescription)
	Prescription := new(Prescription)
    err := db.Preload("User").Find(&Prescription, "user_approved_by = ?", id).Scan(&response).Error
    if err != nil {
        log.Println("error getting prescriptions",err.Error())
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
	role := GetAuthUser(c)
	body := Prescription{}
	body.QuoteNumber = GenerateQuoteNumber()
	prescription := Prescription{
		SubTotal:body.SubTotal,
		QuoteNumber: body.QuoteNumber,
	}

	if err:=c.BodyParser(&body); err != nil{
		return nil,errors.New("failed to parse data")
	}
	prescription.CalculateVAT(vatRate)
	body.SubTotal = prescription.SubTotal
	body.VAT=prescription.VAT
	body.Total=prescription.Total
	body.UserApprovedBy=&user_id
	response := new(ResponsePrescription)

	//find the prescription to be updated
	err := db.First(&prescription,"id = ?",id).Error
	if err != nil {
		log.Println("error updating prescription",err.Error())
		return nil,errors.New("failed to update prescription")
	}
	oldValues := prescription
	//update prescription
	err = db.Model(&prescription).Updates(&body).Scan(&response).Error
	if err != nil{
		log.Println("error updating prescription",err.Error() )
		return nil, errors.New("failed to update prescription")
	}
	newValues := prescription

	//update audit log
	if err := utilities.LogAudit("Update",user_id,role,"Prescription",id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return response,nil
}

/*
Deletes the prescription
@params id
*/
func DeletePrescription(c *fiber.Ctx, id uuid.UUID) error {
	user_id, _:= GetAuthUserID(c)
	role := GetAuthUser(c)
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
    oldValues := prescription
    // Attempt to delete the prescription
    err = db.Delete(&prescription).Error
    if err != nil {
        log.Printf("Error deleting prescription with ID %s: %v", id, err)
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete prescription")
    }
	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"Prescription",id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
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
calculate subtotal
*/
func CalculateSubTotal(prescription *Prescription) float64 {
    var subTotal float64 = 0.0

    // Loop through each QuoteDetail and calculate the total
    for _, quoteDetail := range prescription.QuoteDetail {
        itemTotal := (quoteDetail.Unit * quoteDetail.Quantity) - quoteDetail.Discount + quoteDetail.Vat
        subTotal += itemTotal
    }
	log.Println(subTotal)
    return subTotal
}

/*
VAT calculator method
@params vatRate
*/
func (p *Prescription)CalculateVAT(vatRate float64){
	p.VAT = p.SubTotal*vatRate/100
	p.Total = p.SubTotal+p.VAT
}
// GenerateQuoteNumber is a placeholder for your quote number generation logic
func GenerateQuoteNumber() string {
    datePrefix := time.Now().Format("20060102") // YYYYMMDD format
    randomSuffix := rand.Intn(9999) // Random 4-digit number
    return fmt.Sprintf("QUO-%s-%04d", datePrefix, randomSuffix)
}