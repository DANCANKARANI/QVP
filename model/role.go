package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// adds role
func CreateRole(c *fiber.Ctx, body Role) error {
	id := uuid.New()
	body.ID = id
	err := db.Create(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to add role")
	}
	return nil
}

//gets roles

func GetRoles(c *fiber.Ctx) (*[]Role, error) {
	response := []Role{}
	err := db.Model(&Role{}).Scan(&response).Error
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
func UpdateRole(c *fiber.Ctx, roleID uuid.UUID) (*Role, error) {
	updatedRole := &Role{}

	// Parse request body into a Role struct
	if err := c.BodyParser(updatedRole); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to update role")
	}

	// Find the role by roleID and update it with the new data
	if err := db.First(&Role{},"id = ?", roleID).Updates(updatedRole).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil, errors.New("no roles found for update")
		}
		return nil, errors.New("failed to update role")
	}

	return updatedRole, nil
}

/*
deletes a role
@params role_id
*/
func DeleteRole(c *fiber.Ctx, role_id uuid.UUID) error {
	if err := db.First(&Role{}, "id = ?", role_id).Delete(&Role{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return errors.New("failed to delete role")
		}
		return errors.New("failed to delete role")
	}
	return nil
}
