package model

import (
	"errors"
	"log"
	"strconv"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//adds failed job
func AddFailedJob(c *fiber.Ctx)(*FailedJob,error){
	failedJob := new(FailedJob)

	//get request body
	if err := c.BodyParser(&failedJob); err != nil{
		log.Println("failed to parse request data", err.Error())
		return nil, errors.New("failed to parse request data")
	}
	id :=uuid.New()
	failedJob.ID = id
	//create faildJob
	if err := db.Create(&failedJob).Error; err != nil{
		log.Println("error creating failed job:",err.Error())
		return nil, errors.New("failed add failed job")
	}
	newValues := failedJob

	//update audit log
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	if err := utilities.LogAudit("Create",user_id,role,"Failed Job",id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}

	return newValues, nil
}

/*
updates failed jobs
@params failed_job_id
*/
func UpdateFailedJob(c *fiber.Ctx,failed_job_id uuid.UUID)(*FailedJob, error){
	failedJob := new(FailedJob)

	//get request body
	if err := c.BodyParser(&failedJob); err != nil{
		log.Println("error parsing request body:",err.Error())
		return nil, errors.New("failed to update failed job")
	}

	//get old values
	if err := db.First(failedJob,"id = ?",failed_job_id).Error; err != nil{
		log.Println("error finding failed job to be updated:", err.Error())
		return nil,errors.New("error updating failed job")
	}
	oldValues := failedJob

	//update failed jobs
	if err := db.Model(failedJob).Updates(failedJob).Error; err != nil{
		log.Println("error updating failed job:", err.Error())
		return nil, errors.New("failed to update failed job")
	}
	newValues := failedJob

	//update audit log
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	if err := utilities.LogAudit("Update",user_id,role,"Failed Job",failed_job_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

	//return response
	return newValues, nil
}

/*
deletes failed jobs
@params faile_job_id
*/
func DeleteFaileJob(c *fiber.Ctx, failed_job_id uuid.UUID)(error){
	failedJob := new(FailedJob)

	//get old values
	if err := db.First(&failedJob,"id = ?",failed_job_id).Error; err != nil{
		log.Println("error getting failed job for delition:",err.Error())
		return errors.New("failed to delete failed jobs")
	}
	oldValues := failedJob

	//delete failed job
	if err := db.Delete(&failedJob).Error; err != nil{
		log.Println("error deleting failed jobs:",err.Error())
		return errors.New("failed to delete faild job")
	}

	//update audit log
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)
	if err := utilities.LogAudit("Deleted",user_id,role,"Failed Job",failed_job_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}

	return nil
}

//get all failed jobs
func GetFailedJobs(c *fiber.Ctx)(*[]FailedJob, error){
	//get page
	page, err := strconv.Atoi(c.Query("page"))
	if err !=nil || page < 1 {
		log.Println(err.Error())
		return nil,errors.New("invalid page number")
	}
	//get page size
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize < 1 {
		log.Println(err.Error())
		return nil,errors.New("invalid page size")
	}
	//response user
	failedJob := new([]FailedJob)
	var totalPages int64
	//get prescription and count
	db.Model(&Prescription{}).Count(&totalPages)
	db.Offset((page -1) * pageSize).Limit(pageSize).Find(&failedJob)
	
	//return response
	return failedJob, nil
}