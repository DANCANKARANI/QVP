package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type ResponseQuote struct{
	ID uuid.UUID		`json:"id"`
	PrescriptionId 	uuid.UUID 		`json:"prescription_id"`
	Description		string			`json:"description"`
	Unit			float64			`json:"unit"`
	Quantity		int				`json:"quantity"`
	Measure			string			`json:"measure"`
	Price			float64			`json:"price"`
	Discount 		float64			`json:"discount"`
	Vat				float64			`json:"vat"`
	Total			float64			`json:"total"`
}
//adds a quote detail
func AddQuoteDetail(c *fiber.Ctx) (*ResponseQuote, error) {
	user_id, _:=GetAuthUserID(c)
	role := GetAuthUser(c)

	quoteDetail := new(QuoteDetail)
	response := new(ResponseQuote)
	quoteDetail.ID=uuid.New()
	//parse request body
	if err := c.BodyParser(&quoteDetail); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	//calculate vat
	quoteDetail.CalculateVAT(vatRate)

	quoteDetail.ID=uuid.New()
	newValues := quoteDetail
	//create quote detail
	err := db.Create(&quoteDetail).Scan(&newValues).Scan(&response).Error
	if err!=nil {
		log.Println(err.Error())
		return nil, errors.New("failed to add quote details")
	}
	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"duote details",quoteDetail.ID,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
	//return response
	return response, nil
}

/*
updates quote detail
@params quote_detail_id
*/
func UpdateQuoteDetail(c *fiber.Ctx, quote_detail_id uuid.UUID) (*ResponseQuote,int,error){
	user_id, _:=GetAuthUserID(c)
	role := GetAuthUser(c)
	body := QuoteDetail{}
	//parse request body into quoteDetail
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,fiber.StatusInternalServerError, errors.New("failed to parse json data")
	}
	//find the record
	body.CalculateVAT(vatRate)
	
	//get old values
	quoteDetail := QuoteDetail{}
	if err := db.First(&quoteDetail, "id = ?", quote_detail_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil, fiber.StatusNotFound, errors.New("record not found")
		}
		log.Println(err.Error())
		return nil, fiber.StatusInternalServerError, errors.New("failed to find record")
	}
	//uppdate quote details
	oldValues := quoteDetail
	newValues:=new(ResponseQuote)
	if err := db.Model(&quoteDetail).Updates(body).Scan(&newValues).Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.StatusInternalServerError, errors.New("failed to update quote details")
	}

	//update audit log
	if err := utilities.LogAudit("Update",user_id,role,"quote details",quote_detail_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return newValues,fiber.StatusOK,nil
}
/*
deletes quote details
@params quote_detail_id
*/
func DeleteQuoteDetail(c *fiber.Ctx, quote_detail_id uuid.UUID)(int, error) {
	user_id, _:=GetAuthUserID(c)
	role := GetAuthUser(c)
	quoteDetail := new(QuoteDetail)
	oldValues := quoteDetail
	if err := db.First(&oldValues, "id = ?",quote_detail_id).Delete(&quoteDetail).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return fiber.StatusNotFound, errors.New("record not found")
		}
		log.Println(err.Error())
		return fiber.StatusInternalServerError, errors.New("failed delete quote detail")
	}
	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"quote details",quote_detail_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
	return fiber.StatusOK,nil
}
/*
gets Quote with prescriptions
@params prescription_id
*/
func GetQuoteDetailsWithPrescription(prescription_id uuid.UUID)(int, *[]QuoteDetail,error){
	quoteDetail := new([]QuoteDetail)
	//getting data
	if err := db.Preload("Prescription").Where("prescription_id = ?", prescription_id).Find(&quoteDetail).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return fiber.StatusNotFound,nil,errors.New("record not found")
		}
		log.Println(err.Error())
		return fiber.StatusInternalServerError,nil,errors.New("failed to get quote details")
	}
	return fiber.StatusOK,quoteDetail,nil
}

/*
calculates the vat
@params vatRate
*/
func (q *QuoteDetail) CalculateVAT(vatRate float64){
	q.Price = q.Unit*q.Quantity
	q.Vat = q.Price *vatRate/100
	q.Total = q.Price+q.Vat- q.Discount
}
//func HandleDbError(err error)() 