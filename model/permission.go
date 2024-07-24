package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
adds permissions
*/
func CreatePermission(c *fiber.Ctx) error {
	body:=Permission{}
	id:=uuid.New()
	body.ID = id
	if err := c.BodyParser(&body);err != nil {
		return errors.New("failed to add permission")
	}
	err := db.Create(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to add permission")
	}
	return nil
}
/*
gets permissions
*/
func GetPermission(c *fiber.Ctx) (*[]Permission, error) {
	response := new([]Permission)
	err := db.Model(&Permission{}).Scan(&response).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil, errors.New("no permissions found")
		}
		log.Println(err.Error())
		return nil, errors.New("failed to get permissions")
	}
	return response, nil
}

/*
updates a permission
@params permission_id
*/
func UpdatePermission(c *fiber.Ctx, permissionID uuid.UUID) (*Permission, error) {
	updatedPermission := new(Permission)

	// Parse request body into a permission struct
	if err := c.BodyParser(&updatedPermission); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to update permission")
	}

	// Find the permission by permissionID and update it with the new data
	if err := db.First(&Permission{},"id = ?", permissionID).Updates(updatedPermission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil, errors.New("no permission found for update")
		}
		return nil, errors.New("failed to update permission")
	}

	return updatedPermission, nil
}

/*
deletes a role
@params role_id
*/
func DeletePermission(c *fiber.Ctx, permmission_id uuid.UUID) error {
	if err := db.First(&Permission{}, "id = ?", permmission_id).Delete(&Permission{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return errors.New("failed to delete permission")
		}
		return errors.New("failed to delete permission")
	}
	return nil
}
//creates association permissions
func AssociatePermissions(c *fiber.Ctx)error{
	type Request struct{
		RoleID  		uuid.UUID 		`json:"role_id"`
		PermissionIDs	[]uuid.UUID		`json:"permission_ids"`
	}
	req :=new(Request)
	if err := c.BodyParser(req);err !=nil {
		log.Println(err.Error())
		return errors.New("cannot parse JSON")
	}
	var role Role
	if err :=db.First(&role,"id = ?",req.RoleID).Error; err != nil {
		log.Println(err.Error())
		return errors.New("role not found")
	}
	var permissions []Permission
	if err := db.Where("id IN ?",req.PermissionIDs).Find(&permissions).Error; err != nil {
		log.Println(err.Error())
		return errors.New("failed to find permissions")
	}

	if err := db.Model(&role).Association("Permissions").Append(&permissions);err != nil{
		log.Println(err.Error())
		return errors.New("failed to associate permissions")
	}
	return nil
}