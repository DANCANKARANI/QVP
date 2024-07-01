package model

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SendNotification(c *fiber.Ctx, body Notification ) (*Notification,error){
	db.AutoMigrate(&body)
	body.ID = uuid.New()
	id, _:=GetAuthUserID(c)
	body.UserID=id
	body.CreatedAt = time.Now().Local()
	body.UpdatedAt = time.Now().Local()
	err := db.Create(&body).Error
	if err != nil {
		return nil,errors.New("failed to add notification")
	}
	return &body,nil
}
/*
gets the notification
@params id
@params c *fiber.Ctx
*/
func GetNotification(c *fiber.Ctx, id uuid.UUID)(*Notification,error){
	notification := Notification{}
	err :=db.Model(&notification).Where("user_id = ?",id).Scan(&notification).Error
	if err != nil {
		return nil,errors.New("failed to get notifications:"+err.Error())
	}
	return &notification,nil
}