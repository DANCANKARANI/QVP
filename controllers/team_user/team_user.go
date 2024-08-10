package team_user

import (
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//add user to a team handler
func AddUserToTeamHandler(c *fiber.Ctx)error{
	user_id,_:=uuid.Parse(c.Params("user_id"))
	team_id,_:=uuid.Parse(c.Params("team_id"))
	response, err := model.AddUserToTeam(c,team_id,user_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added user to a team",fiber.StatusOK,response)
}
//remove user from team handler
func RemoveUserFromTeamHandler(c *fiber.Ctx)error{
	user_id,_:=uuid.Parse(c.Params("id"))
	team_id,_:=uuid.Parse(c.Params("id"))
	code, err := model.RemoveUserFromTeam(c,user_id,team_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),code)
	}
	return utilities.ShowMessage(c,"successfully removed user from team",code)
}

//get users from team hander
func GetUsersFromTeamHandler(c *fiber.Ctx)error{
	team_id,_:=uuid.Parse(c.Params("team_id"))
	response, err := model.GetUsersFromTeam(team_id)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved users from team",fiber.StatusOK,response)
}
//get user users' teams handler
func GetTeamsForUserHandler(c *fiber.Ctx)error{
	team_id,_:=uuid.Parse(c.Params("team_id"))
	response, err :=model.GetTeamsForUser(team_id)
	if err != nil{
		return utilities.ShowError(c, err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfull retrieved teams for users", fiber.StatusOK,response)
}
//updates team user
func UpdateUserInTeamHandler(c *fiber.Ctx)error{
	user_id, _:=uuid.Parse(c.Params("user_id"))
	team_id, _:= uuid.Parse(c.Params("team_id"))
	response, err := model.UpdateUsersInTeam(c,user_id,team_id)
	if err != nil{
		return utilities.ShowError(c, err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated users in a team", fiber.StatusOK,response)
}