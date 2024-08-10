package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
adds permissions
*/
func CreatePermission(c *fiber.Ctx) error {
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)

	body:=Permission{}
	id:=uuid.New()
	body.ID = id

	//get request body
	if err := c.BodyParser(&body);err != nil {
		return errors.New("failed to add permission")
	}

	//create permission
	err := db.Create(&body).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to add permission")
	}

	newValues := body
	
	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"Permission",id,nil,newValues,c); err != nil{
		log.Println(err.Error())
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
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	permission := new(Permission)
	updatedPermission := new(Permission)

	// Parse request body into a permission struct
	if err := c.BodyParser(&updatedPermission); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to update permission")
	}

	// Find the permission by permissionID 
	if err := db.First(&permission,"id = ?", permissionID).Error; err != nil {
		log.Println("error updating permission",err.Error())
		return nil, errors.New("failed to update permission")
	}
	oldValues := permission

	//update permission
	if err := db.Model(&permission).Updates(&updatedPermission).Error; err != nil{
		log.Println("errors updating permission",err.Error())
		return nil, errors.New("failed to update permission")
	}
	newValues := permission

	//update audit logs
	if err := utilities.LogAudit("Update",user_id,role,"Permission",permission.ID,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
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