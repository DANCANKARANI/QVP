package role

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add role handler
func AddRoleHandler(c *fiber.Ctx)error{
	body := model.Role{}
	if err := c.BodyParser(&body); err != nil{
		return utilities.ShowError(c,"failed add role",fiber.StatusInternalServerError)
	}
	if err:= model.CreateRole(c,body); err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"role added successfully",fiber.StatusCreated)
}
//get roles handler
func GetRolesHandler(c *fiber.Ctx)error{
	Roles,err:=model.GetRoles(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved roles",fiber.StatusOK,Roles)
}
//update role 
func UpdateRoleHandler(c *fiber.Ctx)error{
	role_id,_:=uuid.Parse(c.Params("id"))
	Role,err := model.UpdateRole(c,role_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated the role",fiber.StatusOK,Role)
}
//delete role
func DeleteRoleHandler(c *fiber.Ctx)error{
	role_id,_:=uuid.Parse(c.Params("id"))
	if err := model.DeleteRole(c,role_id); err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted the role",fiber.StatusOK)
}
//create association handler
func AssociatePermissionsHandler(c *fiber.Ctx)error{
	if err:=model.AssociatePermissions(c);err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully added assoiation permission", fiber.StatusOK)
}