package model

import (
	"errors"
	"log"

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
	quoteDetail := new(QuoteDetail)
	response := new(ResponseQuote)
	quoteDetail.ID=uuid.New()
	if err := c.BodyParser(&quoteDetail); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	quoteDetail.CalculateVAT(vatRate)

	quoteDetail.ID=uuid.New()
	err := db.Create(&quoteDetail).Scan(&response).Error
	if err!=nil {
		log.Println(err.Error())
		return nil, errors.New("failed to add quote details")
	}
	return response, nil
}

/*
updates quote detail
@params quote_detail_id
*/
func UpdateQuoteDetail(c *fiber.Ctx, quote_detail_id uuid.UUID) (*ResponseQuote,int,error){
	body := QuoteDetail{}
	response :=  new(ResponseQuote)
	//parse request body into quoteDetail
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,fiber.StatusInternalServerError, errors.New("failed to parse json data")
	}
	//find the record
	body.CalculateVAT(vatRate)
	if err := db.First(&QuoteDetail{}, "id = ?",quote_detail_id).Updates(&body).Scan(&response).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil, fiber.StatusNotFound,errors.New("record not found")
		}
		log.Println(err.Error())
		return nil, fiber.StatusInternalServerError,errors.New("failed to update quote details")
	}
	return response,fiber.StatusOK,nil
}
/*
deletes quote details
@params quote_detail_id
*/
func DeleteQuoteDetail(c *fiber.Ctx, quote_detail_id uuid.UUID)(int, error) {
	quoteDetail := new(QuoteDetail)
	if err := db.First(&quoteDetail, "id = ?",quote_detail_id).Delete(&quoteDetail).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return fiber.StatusNotFound, errors.New("record not found")
		}
		log.Println(err.Error())
		return fiber.StatusInternalServerError, errors.New("failed delete quote detail")
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