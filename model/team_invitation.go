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
	user_id, _ := GetAuthUserID(c)
	role :=GetAuthUser(c)
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
	//update audit logs
	newValues := teamInvitation
	if err := utilities.LogAudit("Create",user_id,role,"Team invitation",user_id,nil,newValues,c); err != nil{
		log.Println(err.Error())
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
	role := GetAuthUser(c)
	user_id, _ :=GetAuthUserID(c)
	teamInvitation := new(TeamInvitation)
	response := new(ResponseTeamInvitation)
	if err := c.BodyParser(&teamInvitation);err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	//check existence
	err := db.First(&teamInvitation,"id= ?",id).Scan(&teamInvitation).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("record not found")
	}
	oldValues :=teamInvitation
	//update the invatitation teams
	err = db.Model(&teamInvitation).Where("id = ?",id).Updates(&teamInvitation).Scan(&response).Error
	if err !=nil{
		log.Println(err.Error())
		return nil,errors.New("failed to find team invitation")
	}
	//update audit log
	newValues := teamInvitation
	if err := utilities.LogAudit("Create",user_id,role,"Team invitation",user_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	
	return response,nil
}

/*
deletes the invitation team
@params invitation_team_id
*/
func DeleteInvitationTeam(c *fiber.Ctx,id uuid.UUID)(error){
	role := GetAuthUser(c)
	user_id,_:=GetAuthUserID(c)
	teamInvitation :=new(TeamInvitation)

	//find the team invitation
	err :=db.First(&teamInvitation, "id = ?", id).Scan(&teamInvitation).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete")
	}

	oldValues := teamInvitation
	//delete team invitations
	err = db.Model(&teamInvitation).Delete(&teamInvitation).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete")
	}

	//update audit logs
	if err := utilities.LogAudit("Delete",user_id,role,"Team invitation",user_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}
	return nil
}