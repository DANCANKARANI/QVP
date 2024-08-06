package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
adds user to a team
@params user_id, team_id
*/
func AddUserToTeam(c *fiber.Ctx, team_id, user_id uuid.UUID) (*TeamUser, error) {
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
	return &teamUser, nil
}

/*
remove user in a team
@params team_id, user_id
*/
func RemoveUserFromTeam(team_id,user_id uuid.UUID)(code int,err error){
	teamUser := new(TeamUser)
	err =db.Where("user_id = ? AND team_id =?",user_id,team_id).First(teamUser).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println("record not found")
			return fiber.StatusNotFound,errors.New("record not found")
		}
		log.Println("database error:",err.Error())
		return fiber.StatusInternalServerError, errors.New("failed to remove user from team")
	}
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
	teamUser := new(TeamUser)
	if err:= c.BodyParser(&teamUser); err != nil{
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}
	err := db.Where("user_id = ? AND team_id =?",user_id, team_id).First(&teamUser).Updates(&teamUser).Scan(teamUser).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil, errors.New("record not found")
		}
		log.Println(err.Error())
		return nil, errors.New("failed to update user in a team")
	}
	return teamUser, nil
}
