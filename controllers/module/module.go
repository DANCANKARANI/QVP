package module

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//create module handler
func CreateModuleHandler(c *fiber.Ctx)error{
	module,err := model.CreateModule(c) 
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"module added successfully",fiber.StatusOK,module)
}
//get modules handler
func GetModulesHandler(c *fiber.Ctx)error{
	module,err := model.GetModules()
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved modules",fiber.StatusOK,module)
}
//update module handler
func UpdateModuleHandler(c *fiber.Ctx)error{
	module_id,_:=uuid.Parse(c.Params("id"))
	module,err := model.UpdateModule(c,module_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated the module",fiber.StatusOK,module)
}

//delete module handler
func DeleteModuleHandler(c *fiber.Ctx)error{
	module_id,_:=uuid.Parse(c.Params("id"))
	if err := model.DeleteModule(c,module_id); err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"module deleted successfully",fiber.StatusOK)
}
