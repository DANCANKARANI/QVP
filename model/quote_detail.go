package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
//adds a quote detail
func AddQuoteDetail(c *fiber.Ctx) (*QuoteDetail, error) {
	quoteDetail := new(QuoteDetail)
	quoteDetail.ID=uuid.New()
	if err := c.BodyParser(&quoteDetail); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	quoteDetail.CalculateVAT(vatRate)

	quoteDetail.ID=uuid.New()
	err := db.Create(&quoteDetail).Error
	if err!=nil {
		log.Println(err.Error())
		return nil, errors.New("failed to add quote details")
	}
	return quoteDetail, nil
}

/*
updates quote detail
@params quote_detail_id
*/
func UpdateQuoteDetail(c *fiber.Ctx, quote_detail_id uuid.UUID) (*QuoteDetail,int,error){
	quoteDetail := new(QuoteDetail)
	if err := c.BodyParser(&quoteDetail); err != nil {
		log.Println(err.Error())
		return nil,fiber.StatusInternalServerError, errors.New("failed to parse json data")
	}
	if err := db.First(&quoteDetail, "id = ?",quote_detail_id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil, fiber.StatusNotFound,errors.New("record not found")
		}
		log.Println(err.Error())
		return nil, fiber.StatusInternalServerError,errors.New("failed to update quote details")
	}
	if err:=db.Model(&quoteDetail).Updates(&quoteDetail). Error; err != nil{
		log.Println(err.Error())
		return nil, fiber.StatusInternalServerError,errors.New("failed to update quote details")
	}
	return quoteDetail,fiber.StatusOK,nil
}

/*
deletes quote details
@params quote_detail_id
*/
func DeleteQuoteDetail(c *fiber.Ctx, quote_detail_id uuid.UUID)(int, error) {
	quoteDetail := new(QuoteDetail)
	if err := db.First(&quoteDetail, "id = ?",quote_detail_id).Error; err != nil{
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
	if err := db.Where("prescription_id = ?", prescription_id).Find(&quoteDetail).Error; err != nil{
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