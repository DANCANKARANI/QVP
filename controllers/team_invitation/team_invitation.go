package team_invitation

import (
	"log"

	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add team invitation handler
func AddTeamInvitationHandler(c *fiber.Ctx) error {
	teamInvitation, err := model.CreateTeamInvitation(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added team invitation",fiber.StatusOK,teamInvitation)
}
//get team invitations handler
func GetTeamInvitationsHandler(c *fiber.Ctx)error{
	teamInvitation,err :=model.GetTeamInvitations()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved team invitations",fiber.StatusOK,teamInvitation)
}

//get team invitation by team handler
func GetTeamInvitationByTeamHandler(c *fiber.Ctx)error{
	id,err:=uuid.Parse(c.Params("id"))
	if err != nil{
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to get team invitation",fiber.StatusNotFound)
	}
	teamInvitation,err :=model.GetTeamInvitationByTeam(id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved team invitations",fiber.StatusOK,teamInvitation)
}

//update team invitation handler
func UpdateTeamInvitationHandler(c *fiber.Ctx)error{
	id,err:=uuid.Parse(c.Params("id"))
	if err != nil{
		log.Println(err.Error())
		return utilities.ShowError(c,"failed update team invitation",fiber.StatusNotFound)
	}
	teamInvitation,err:=model.UpdateTeamInvitation(c,id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated team invitation",fiber.StatusOK,teamInvitation)
}

//delete team invitation handler
func DeleteTeamInvitationHandler(c *fiber.Ctx)error{
	id,err:=uuid.Parse(c.Params("id"))
	if err != nil{
		log.Println(err.Error())
		return utilities.ShowError(c,"failed delete team invitation",fiber.StatusNotFound)
	}
	if err :=model.DeleteInvitationTeam(id); err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}	
	return utilities.ShowMessage(c,"successfully deleted team invitation",fiber.StatusOK)
}