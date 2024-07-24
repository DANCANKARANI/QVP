package permission

import (
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
//add permission handler
func AddPermissionHandler(c *fiber.Ctx)error{
	if err:= model.CreatePermission(c); err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"permission added successfully",fiber.StatusCreated)
}
//get permission handler
func GetPermissionHandler(c *fiber.Ctx)error{
	Roles,err:=model.GetPermission(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved permissions",fiber.StatusOK,Roles)
}
//update permission handler 
func UpdatePermissionHandler(c *fiber.Ctx)error{
	role_id,_:=uuid.Parse(c.Params("id"))
	Role,err := model.UpdatePermission(c,role_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated permission",fiber.StatusOK,Role)
}
//delete permission handler
func DeletePermissionHandler(c *fiber.Ctx)error{
	permission_id,err:=uuid.Parse(c.Params("id"))
	if err!=nil{
		log.Println(err.Error())
	}
	if err := model.DeletePermission(c,permission_id); err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted permission",fiber.StatusOK)
}

