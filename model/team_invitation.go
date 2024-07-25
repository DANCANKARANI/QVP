package model

import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
func GetTeamInvitations()(*[]TeamInvitation, error){
	teamInvitation := new([]TeamInvitation)
	err:=db.Model(&teamInvitation).Preload("Teams").Scan(&teamInvitation).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil,errors.New("record not found")
		}
		log.Println(err.Error())
		return nil,errors.New("failed to get team invitations")
	}
	return teamInvitation,nil
}
/*
gets teamsInvitations by team
@params team_id
*/
func GetTeamInvitationByTeam(team_id uuid.UUID)(*TeamInvitation,error){
	team :=new(TeamInvitation)
	err :=db.Where("team_id = ?", team_id).Find(&team).Error
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil, errors.New("record not found")
		}
		log.Println(err.Error())
		return nil, errors.New("failed to get teams")
	}
	return team,nil
} 

/*
updates teamInvitations
@params team_invitation_id
*/
func UpdateTeamInvitation(c *fiber.Ctx, id uuid.UUID)(*TeamInvitation, error) {
	teamInvitation := new(TeamInvitation)
	if err := c.BodyParser(&teamInvitation);err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	//check existence
	err := db.First(&teamInvitation,"id = ?",id).Error
	if err !=nil{
		log.Println(err.Error())
		return nil,errors.New("failed to get team")
	}
	//update the invatitation teams
	err =db.Updates(&teamInvitation).Scan(&teamInvitation).Error
	if err !=nil{
		log.Println(err.Error())
		return nil,errors.New("failed to update")
	}
	return teamInvitation,nil
}
/*
deletes the invitation team
@params invitation_team_id
*/
func DeleteInvitationTeam(id uuid.UUID)(error){
	teamInvitation :=new(TeamInvitation)
	err :=db.First(&teamInvitation, "id = ?", id).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete")
	}
	err = db.Delete(&teamInvitation).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete")
	}
	return nil
}

