package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//adds new branch
func CreateBranch(c *fiber.Ctx)(*Branch,error) {
	id := uuid.New()
	branch := new(Branch)
	//parse requset body
	if err := c.BodyParser(&branch); err != nil{
		log.Println("error parsing request body",err.Error())
		return nil, errors.New("failed to parse request body")
	}

	//create branch
	branch.ID = id
	if err := db.Create(&branch).Error; err != nil{
		log.Println("error adding branch")
		return nil, errors.New("failed to add branch")
	}
	newValues := branch

	//update audit logs
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	if err := utilities.LogAudit("Create",user_id,role,"Branch",id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return newValues, nil
}

/*
updates branch
@params branch_id
*/
func UpdateBranch(c *fiber.Ctx,branch_id uuid.UUID)(*Branch, error){
	branch := new(Branch)

	//parse request body
	if err := c.BodyParser(&branch); err != nil{
		log.Println("error parsing branch request body",err.Error())
		return nil, errors.New("failed to parse request body")
	}

	//get old values
	if err := db.First(branch,"id = ?",branch_id).Error; err != nil{
		log.Println("error getting branch for update", err.Error())
		return nil, errors.New("failed to update branch")
	}
	oldValues := branch

	//udpate branch
	if err := db.Model(branch).Updates(branch).Error; err != nil{
		log.Println("error updating branch", err.Error())
		return nil, errors.New("failed to update branch")
	}
	newValues := branch

	//update audit logs
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	if err := utilities.LogAudit("Update",user_id,role,"Branch",branch_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

	//return response
	return newValues, nil
}

/*
deletes branch
@params branch_id
*/
func DeleteBranch(c *fiber.Ctx, branch_id uuid.UUID)(error){
	branch := new(Branch)

	//get oldValues
	if err := db.First(branch,"id =?",branch_id).Error; err != nil{
		log.Println("error finding branch for delition",err.Error())
		return errors.New("failed to delete branch")
	}
	oldValues := branch

	//delete branch 
	if err := db.Delete(&branch).Error; err != nil{
		log.Println("error deleting branch", err.Error())
		return errors.New("failed to delete branch")
	}

	//update audit logs
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	if err := utilities.LogAudit("Delete",user_id,role,"Branch",branch_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}

	return nil
}

//get all branches
func GetBranch(c *fiber.Ctx)(*[]Branch, error){
	branch := new([]Branch)

	//get branches
	if err := db.Find(&branch).Error; err != nil{
		log.Println("error getting branches",err.Error())
		return nil,errors.New("failed to get branches")
	}

	//return response
	return branch, nil
}
