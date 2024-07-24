package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//creates a module
func CreateModule(c *fiber.Ctx) (*Module,error) {
	id := uuid.New()
	module := new(Module)
	if err := c.BodyParser(module); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to parse json data")
	}
	module.ID = id 
	if err := db.Create(&module).Error; err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add module")
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
    var existingModule Module
    if err := db.First(&existingModule, "id = ?", moduleID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println("record not found")
            return nil, errors.New("record not found")
        }
        log.Println(err.Error())
        return nil, errors.New("failed to find module")
    }

    updatedModule := new(Module)
    if err := c.BodyParser(updatedModule); err != nil {
        log.Println(err.Error())
        return nil, errors.New("failed to parse json data")
    }

    if err := db.Model(&existingModule).Updates(updatedModule).Error; err != nil {
        log.Println(err.Error())
        return nil, errors.New("failed to update module")
    }

    return &existingModule, nil
}
/*
deletes the module
@params module_id
*/
func DeleteModule(moduleID uuid.UUID) error {
    module := new(Module)
    if err := db.First(&module, "id = ?", moduleID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println("record not found")
            return errors.New("record not found")
        }
        log.Println(err.Error())
        return errors.New("failed to find module")
    }

    if err := db.Delete(&module).Error; err != nil {
        log.Println(err.Error())
        return errors.New("failed to delete module")
    }
    return nil
}