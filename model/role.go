package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// adds role
func CreateRole(c *fiber.Ctx, body Role) error {
	user_id, _:=GetAuthUserID(c)
	role := GetAuthUser(c)

	id := uuid.New()
	body.ID = id
	err := db.Create(&body).Scan(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to add role")
	}
	newValues := body
	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"role",id,nil,newValues,c); err != nil{
		log.Println(err.Error())
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
	user_id,_:=GetAuthUserID(c)
	role := GetAuthUser(c)
	updatedRole := &Role{}

	// Parse request body into a Role struct
	if err := c.BodyParser(updatedRole); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to update role")
	}

	// Find the role by roleID and update it with the new data
	oldValues:=Role{}
	if err := db.First(&oldValues,"id = ?", roleID).Updates(updatedRole).Scan(&updatedRole).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil, errors.New("no roles found for update")
		}
		return nil, errors.New("failed to update role")
	}

	//update audit logs
	newValues := updatedRole
	if err := utilities.LogAudit("Update",user_id,role,"Role",roleID,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
	//response
	return updatedRole, nil
}

/*
deletes a role
@params role_id
*/
func DeleteRole(c *fiber.Ctx, role_id uuid.UUID) error {
	role := GetAuthUser(c)
	user_id, _:=GetAuthUserID(c)
	oldValues := Role{}

	//find and delete role
	if err := db.First(&oldValues, "id = ?", role_id).Delete(&Role{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return errors.New("failed to delete role")
		}
		return errors.New("failed to delete role")
	}
	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"Role",role_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
	//response
	return nil
}
