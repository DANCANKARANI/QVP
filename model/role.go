package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//adds role
func CreateRole(c *fiber.Ctx, body Role)error{
	err := db.Create(&body).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to add role")
	}
	return nil
}

//gets roles

func GetRoles(c *fiber.Ctx) (*Role, error) {
	response := Role{}
	err := db.Model(&response).Scan(&response).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil, errors.New("no roles found")
		}
		log.Println(err.Error())
		return nil, errors.New("failed to get roles")
	}
	return &response, nil
}

/*
updates a role
@params role_id
*/
func UpdateRoles(c *fiber.Ctx,role_id uuid.UUID)(*Role,error){
	body := Role{}
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to update role")
	}
	err := db.Model(&body).First(&body, "id = ?", role_id).Scan(&body).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil, errors.New("no roles found for update")
		}
		return nil, errors.New("failed to update role")
	}
	return &body,nil
}
/*
deletes a role
@params role_id
*/
func DeleteRole(c *fiber.Ctx,role_id uuid.UUID)error{
	if err:= db.First(&Role{},"id = ?",role_id).Delete(&Role{}).Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return errors.New("failed to delete role")
		}
		return errors.New("failed to delete role")
	}
	return nil
}
