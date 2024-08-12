package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//geolocation details
type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//delivery details
type DeliveryDetail struct {
	Geolocation    Geolocation `json:"geolocation"`
	PickupLocation string      `json:"pickup_location"`
}
// Scan implements the sql.Scanner interface to read from a database value
func (d *DeliveryDetail) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert database value to byte slice")
	}

	if err := json.Unmarshal(bytes, d); err != nil {
		return fmt.Errorf("failed to unmarshal JSON to DeliveryDetails: %w", err)
	}

	return nil
}

// Value implements the driver.Valuer interface to write to a database value
func (d DeliveryDetail) Value() (driver.Value, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DeliveryDetails to JSON: %w", err)
	}

	return string(bytes), nil
}
func UpdateDeliveryDetails(c *fiber.Ctx,prescription_id uuid.UUID)(*DeliveryDetail,error) {
	var prescription Prescription

	//parse request body
	var deliveryDetails DeliveryDetail
	if err := c.BodyParser(&deliveryDetails); err != nil{
		log.Println("failed to parse delivery details:",err.Error())
		return nil, errors.New("failed to parse delivery details")
	}

	//find the prescriptio
	if err:=db.First(&prescription,"id = ?",prescription_id).Error; err != nil{
		log.Println("error finding prescription:",err.Error())
		return nil, errors.New("prescription not found")
	}

	//update delivery details in prescription
	prescription.DeliveryDetails= deliveryDetails
	if err := db.Save(&prescription).Error; err != nil{
		log.Println("error updating delivery details:",err.Error())
		return nil, errors.New("failed to update delivery details")
	}

	return &prescription.DeliveryDetails,nil
}