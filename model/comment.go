package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add comment
func CreateComment(c *fiber.Ctx)(*Comment, error){

	comment := new(Comment)
	if err := c.BodyParser(&comment); err != nil{
		log.Println("failed to parse request body:",err.Error())
		return nil, errors.New("failed to parse request body")
	}
	
	comment.ID = uuid.New()
	err := db.Create(&comment).Error
	if err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to create comment")
	}
	return comment, nil
}
/*
updates comment
@params comment_id
*/
func UpdateComment(c *fiber.Ctx, comment_id uuid.UUID)(*Comment, error){
	comment := new(Comment)
	if err := c.BodyParser(&comment); err != nil{
		log.Println("failed to parse request data", err.Error())
		return nil, errors.New("failed to parse request data")
	}

	if err := db.Model(&comment).Where("id = ?",comment_id).Updates(&comment).Scan(&comment).Error; err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to update comments")
	}
	return comment,nil
}

/*
deletes comment
@params comment_id
*/
func DeleteComment(comment_id uuid.UUID)error{
	comment := new(Comment)
	if err := db.First(comment, "id = ?",comment_id).Delete(&comment).Error; err !=nil{
		return  errors.New("failed to delete comment")
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
