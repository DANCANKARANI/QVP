package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add comment
func CreateComment(c *fiber.Ctx)(*Comment, error){
	user_id, _ := GetAuthUserID(c)
	role := GetAuthUser(c)

	comment := new(Comment)

	//parse requst body
	if err := c.BodyParser(&comment); err != nil{
		log.Println("failed to parse request body:",err.Error())
		return nil, errors.New("failed to parse request body")
	}
	
	comment.ID = uuid.New()

	//create comment
	err := db.Create(&comment).Error
	if err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to create comment")
	}
	newValues := comment

	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"Comment",comment.ID,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return comment, nil
}
/*
updates comment
@params comment_id
*/
func UpdateComment(c *fiber.Ctx, comment_id uuid.UUID)(*Comment, error){
	user_id,_:= GetAuthUserID(c)
	role := GetAuthUser(c)
	comment := new(Comment)
	if err := c.BodyParser(&comment); err != nil{
		log.Println("failed to parse request data", err.Error())
		return nil, errors.New("failed to parse request data")
	}

	//find old values
	if err := db.First(&comment,"id = ?",comment_id).Error; err != nil{
		log.Println("error getting comment for update",err.Error())
		return nil, errors.New("failed to update comment")
	}
	oldValues := comment

	//update comment
	if err := db.Model(&comment).Updates(&comment).Error; err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to update comments")
	}
	newValues := comment

	//update audit logs
	if err := utilities.LogAudit("Update",user_id,role,"Comment",comment.ID,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}

	//return response
	return comment,nil
}

/*
deletes comment
@params comment_id
*/
func DeleteComment(c *fiber.Ctx,comment_id uuid.UUID)error{
	user_id,_:=GetAuthUserID(c)
	role := GetAuthUser(c)

	comment := new(Comment)

	//find old values of the comment
	if err := db.First(comment, "id = ?",comment_id).Error; err != nil{
		log.Println("failed to get comment for delition",err.Error())
		return errors.New("failed to delete comment")
	}
	oldValues := comment

	//delete comment
	if err:=db.Delete(&comment).Error; err !=nil{
		log.Println("error deleting the comment")
		return  errors.New("failed to delete comment")
	}

	//update audit logs
	if err := utilities.LogAudit("Create",user_id,role,"Comment",comment.ID,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
	
	return nil
}

//get comments for a specific entity
func GetEntityComment(commentable_type string, commentable_id uuid.UUID)(*[]Comment,error){
	comment := new([]Comment)
	err := db.Model(&comment).
	Where("commentable_id =? AND commentable_type = ?",commentable_id,commentable_type).
	Find(&comment).Error
	if err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to get entities comments")
	}
	return comment,nil
}
