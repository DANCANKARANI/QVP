package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
//creates new team invitations
func CreateTeamInvitation(c *fiber.Ctx) (*TeamInvitation,error){
	teamInvitation := new(TeamInvitation)
	if err := c.BodyParser(&teamInvitation); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add json data")
	}
	teamInvitation.ID=uuid.New()
	err := db.Create(&teamInvitation).Error; 
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add team invitations")
	}
	return teamInvitation,nil
}

//gets all invitations
