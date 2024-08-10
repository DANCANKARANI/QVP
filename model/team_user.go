package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
adds user to a team
@params user_id, team_id
*/
func AddUserToTeam(c *fiber.Ctx, team_id, user_id uuid.UUID) (*TeamUser, error) {
	//get user role
	role := GetAuthUser(c)
	teamUser := TeamUser{}
	if err :=c.BodyParser(&teamUser); err != nil{
		log.Println("error parsing json data:",err.Error())
		return nil,errors.New("failed to add json data")
	}
	teamUser.ID = uuid.New()
	teamUser.UserID = user_id
	teamUser.TeamID = team_id
	err :=db.Create(&teamUser).Error
	if err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to add user to a team")
	}
	newValues := teamUser

	//update log audit
	if err := utilities.LogAudit("Create",user_id,role,"Team",user_id,nil,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return &teamUser, nil
}

/*
remove user in a team
@params team_id, user_id
*/
func RemoveUserFromTeam(c *fiber.Ctx,team_id,user_id uuid.UUID)(code int,err error){
	role := GetAuthUser(c)
	teamUser := new(TeamUser)
	err =db.Where("user_id = ? AND team_id =?",user_id,team_id).First(teamUser).Scan(teamUser).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println("record not found")
			return fiber.StatusNotFound,errors.New("record not found")
		}
		log.Println("database error:",err.Error())
		return fiber.StatusInternalServerError, errors.New("failed to remove user from team")
	}
	//delete team user
	oldValues := teamUser
	err = db.Model(&teamUser).Delete(&teamUser).Error
	if err != nil{
		log.Println(err.Error())
		return fiber.StatusInternalServerError, errors.New("failed to remove user from team")
	}
	
	//update log audit
	if err := utilities.LogAudit("Create",user_id,role,"Team user",user_id,oldValues,nil,c); err != nil{
		log.Println(err.Error())
	}

	//return
	return fiber.StatusOK, nil
}

/*
gets users in team
@params team_id
*/
func GetUsersFromTeam(team_id uuid.UUID) (*[]User, error) {
    var team Team


    err :=db.Preload("User").First(&team,"id = ?",team_id).Error

    if err != nil {
        log.Println("database error:", err)
        return nil, errors.New("failed to get users in the team")
    }

    if len(team.User) == 0 {
        return nil, errors.New("no users found in the specified team")
    }

    return &team.User, nil
}



/*
gets team for users
@params team_id
*/
func GetTeamsForUser(team_id uuid.UUID)(*Team,error){
	team :=new(Team)
	teamUser := new(TeamUser)
	err := db.Where("team_id=?",team_id).Preload("Team").Find(teamUser).Scan(team).Error
	if err != nil{
		log.Println("database error:",err.Error())
		return nil, errors.New("failed to get teams for a user")
	}
	return team,nil
}

/*
update users in a team
@params team_id, user_id
*/
func UpdateUsersInTeam(c *fiber.Ctx,user_id,team_id uuid.UUID)(*TeamUser,error){
	role := GetAuthUser(c)
	teamUser := new(TeamUser)
	if err:= c.BodyParser(&teamUser); err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	//get old values
	err := db.Where("user_id = ? AND team_id =?",user_id, team_id).First(&teamUser).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil, errors.New("record not found")
		}
		log.Println(err.Error())
		return nil, errors.New("failed to update user in a team")
	}
	oldValues := teamUser

	//update
	if err:=db.Where("user_id = ? AND team_id =?",user_id, team_id).Model(&teamUser).Updates(&teamUser).Scan(teamUser).Error; err != nil{
		log.Println("error updating users in a team"+err.Error())
		return nil, errors.New("failed to update user in team")
	}
	newValues := teamUser
	
		//update log audit
	if err := utilities.LogAudit("Update",user_id,role,"Team user",user_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	//response 
	return teamUser, nil
}
