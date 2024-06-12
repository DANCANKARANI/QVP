package model
import (

	"main.go/utilities"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber/v2"
)
func UserExist(c *fiber.Ctx,db *gorm.DB,phone_number string)(interface{},error){
	 existingUser := User{}
	result := db.Where("phone_number = ?",phone_number).First(&existingUser)
	if result.Error != nil {
		//user not found
		utilities.ShowError(c,result.Error.Error(),fiber.StatusInternalServerError)
	}
	return User{}, nil
}