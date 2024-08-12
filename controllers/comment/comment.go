package comment

import (
	"errors"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//create comment handler
func CreateCommentHandler(c *fiber.Ctx)error{
	response, err := model.CreateComment(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"comment added successfully",fiber.StatusOK,response)
}
//update comment handler
func UpdateCommentHandler(c *fiber.Ctx)error{
	comment_id, _ := uuid.Parse(c.Params("id"))

	response, err := model.UpdateComment(c,comment_id)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return utilities.ShowMessage(c,"record not found", fiber.StatusNotFound)
		}
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"comment updated successfully",fiber.StatusOK,response)
}
//deletes comment handler
func DeleteCommentHandler(c *fiber.Ctx)error{
	comment_id, _ := uuid.Parse(c.Params("id"))

	err := model.DeleteComment(c,comment_id)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return utilities.ShowMessage(c,"record not found",fiber.StatusNotFound)
		}
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"comment deleted successfully",fiber.StatusOK)
}

//gets comment handler
func GetEntityCommentHandler(c *fiber.Ctx)error{
	commentale_id,_:=uuid.Parse(c.Query("commentable_id"))
	commentable_type:= c.Query("commentable_type")
	response, err := model.GetEntityComment(commentable_type,commentale_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	msg_str:="successfully retrieved "+commentable_type+" comments"
	return utilities.ShowSuccess(c,msg_str,fiber.StatusOK,response)
}