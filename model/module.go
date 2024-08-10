package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//creates a module
func CreateModule(c *fiber.Ctx) (*Module,error) {
    user_id,_ := GetAuthUserID(c)

    role :=GetAuthUser(c)

	id := uuid.New()

	module := new(Module)

    //get request body
	if err := c.BodyParser(module); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to parse json data")
	}

	module.ID = id 

    //create module
	if err := db.Create(&module).Error; err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add module")
	}
    newValues := module

    //update audit logs
    if err := utilities.LogAudit("Create",user_id,role,"module",id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return  module,nil
}

//gets all the modules
func GetModules()(*[]Module,error){
	module:=new([]Module)

	err := db.Model(&module).Scan(module).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err.Error())
			return nil,errors.New("record not found")
		}
		return nil,errors.New("failed to get modules")
	}

	return module,nil
}
/*
updates the module 
@params module_id
*/
func UpdateModule(c *fiber.Ctx, moduleID uuid.UUID) (*Module, error) {

    user_id, _:= GetAuthUserID(c)

    role := GetAuthUser(c)

    existingModule :=new(Module)

    //find module record
    if err := db.First(&existingModule, "id = ?", moduleID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println("record not found")
            return nil, errors.New("record not found")
        }
        log.Println(err.Error())
        return nil, errors.New("failed to find module")
    }
    oldValues := existingModule

    updatedModule := new(Module)

    //get request body
    if err := c.BodyParser(updatedModule); err != nil {
        log.Println(err.Error())
        return nil, errors.New("failed to parse json data")
    }

    //update module
    if err := db.Model(&existingModule).Updates(updatedModule).Error; err != nil {
        log.Println(err.Error())
        return nil, errors.New("failed to update module")
    }
    newValues :=existingModule

     //update audit logs
     if err := utilities.LogAudit("Update",user_id,role,"module",moduleID,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
    
    //response
    return newValues, nil
}
/*
deletes the module
@params module_id
*/
func DeleteModule(c *fiber.Ctx,moduleID uuid.UUID) error {
    user_id, _ := GetAuthUserID(c)

    role := GetAuthUser(c)

    module := new(Module)
    //find module record
    if err := db.First(&module, "id = ?", moduleID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println("record not found")
            return errors.New("record not found")
        }
        log.Println(err.Error())
        return errors.New("failed to find module")
    }
    oldValues :=module

    //delete module
    if err := db.Delete(&module).Error; err != nil {
        log.Println(err.Error())
        return errors.New("failed to delete module")
    }

     //update audit logs
     if err := utilities.LogAudit("Delete",user_id,role,"module",moduleID,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
    
    return nil
}