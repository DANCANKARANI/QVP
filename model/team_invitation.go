package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type ResponseTeamInvitation struct{
	ID              uuid.UUID           `json:"id"`
    TeamID          uuid.UUID           `json:"team_id"`
    Email           string              `json:"email"`
    Role            string              `json:"role"`
}
//creates new team invitations
func CreateTeamInvitation(c *fiber.Ctx) (*ResponseTeamInvitation,error){
	teamInvitation := TeamInvitation{}
	response := new(ResponseTeamInvitation)
	if err := c.BodyParser(&teamInvitation); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add json data")
	}
	teamInvitation.ID=uuid.New()
	//verify email adress
	email,err:=utilities.ValidateEmail(teamInvitation.Email)
	if err != nil {
		return nil,errors.New(err.Error())
	}
	teamInvitation.Email = *email
	//add team invitation
	err = db.Create(&teamInvitation).Scan(response).Error; 
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add team invitations")
	}
	return response,nil
}
//gets all invitations
func GetTeamInvitations()(*[]ResponseTeamInvitation, error){
	teamInvitation := new([]TeamInvitation)
	response:=new([]ResponseTeamInvitation)
	err:=db.Model(&teamInvitation).Preload("Teams").Scan(&response).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil,errors.New("record not found")
		}
		log.Println(err.Error())
		return nil,errors.New("failed to get team invitations")
	}
	return response,nil
}
/*
gets teamsInvitations by team
@params team_id
*/
func GetTeamInvitationByTeam(team_id uuid.UUID)(*Team,error){
	team :=new(Team)
	response :=new(Team)
	err :=db.Where("id = ?", team_id).Find(&team).Scan(&response).Error
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil, errors.New("record not found")
		}
		log.Println(err.Error())
		return nil, errors.New("failed to get teams")
	}
	return response,nil
} 

/*
updates teamInvitations
@params team_invitation_id
*/
func UpdateTeamInvitation(c *fiber.Ctx, id uuid.UUID)(*ResponseTeamInvitation, error) {
	teamInvitation := new(TeamInvitation)
	response := new(ResponseTeamInvitation)
	if err := c.BodyParser(&teamInvitation);err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	//check existence
	err := db.Model(&teamInvitation).Where("id = ?",id).Updates(&teamInvitation).Scan(&response).Error
	if err !=nil{
		log.Println(err.Error())
		return nil,errors.New("failed to find team invitation")
	}
	//update the invatitation teams
	return response,nil
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
	err = db.Model(&teamInvitation).Delete(&teamInvitation).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete")
	}
	return nil
}