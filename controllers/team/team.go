package team

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//create team handler
func CreateTeamHandler(c *fiber.Ctx) error {
	team,err:=model.CreateTeam(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added team",fiber.StatusOK,team)
}
//get teams handler
func GetTeamsHandler(c *fiber.Ctx)error{
	team,err := model.GetTeams()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"teams retrieved successfully",fiber.StatusOK,team)
}
//update team handler
func UpdateTeamHandler(c *fiber.Ctx)error{
	team_id,_:= uuid.Parse(c.Params("id"))
	team,err :=model.UpdateTeam(c,team_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated team",fiber.StatusOK,team)
}

//delete team handler
func DeleteTeamHandler(c *fiber.Ctx)error{
	team_id,_:=uuid.Parse("id")
	err:=model.DeleteTeam(team_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"deleted team",fiber.StatusOK)
}