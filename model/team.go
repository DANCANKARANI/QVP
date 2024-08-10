package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type ResponseTeam struct{
	
}
//creates a new team
func CreateTeam(c *fiber.Ctx) (*Team, error) {
	//get user_id
	role :=GetAuthUser(c)
	user_id, _:=GetAuthUserID(c)
	team := new(Team)
	if err := c.BodyParser(&team); err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to parse json data")
	}

	team.ID = uuid.New()
	oldValues := ""

	err := db.Create(&team).Error
	if err != nil {
		return nil, errors.New("failed to create team")
	}
	newValues := team
	//update audit log
	if err = utilities.LogAudit("Create",user_id,role,"Team",team.ID,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
 	return team, err
}
//get teams
func GetTeams()(*[]Team, error){
	team:=new([]Team)
	err := db.Model(&team).Scan(&team).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil, errors.New("record not found")
		}
		log.Println(err.Error())
		return nil,errors.New("failed to get teams")
	}
	return team,nil
}
/*
update teams
@params team_id
*/
func UpdateTeam(c *fiber.Ctx, teamID uuid.UUID) (*Team, error) {
	//get user_id
	role:=GetAuthUser(c)
	user_id, _:= GetAuthUserID(c)
    // Parse the request body into a new team struct
    updatedTeam := new(Team)
    if err := c.BodyParser(updatedTeam); err != nil {
        log.Println(err.Error())
        return nil, errors.New("failed to parse json data")
    }

    // Find the existing team in the database
    var existingTeam Team
    if err := db.First(&existingTeam, "id = ?", teamID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println(err.Error())
            return nil, errors.New("record not found")
        }
        log.Println(err.Error())
        return nil, errors.New("failed to find team")
    }

	oldValues := existingTeam
    // Update the existing team with the new values
    if err := db.Model(&existingTeam).Updates(updatedTeam).Scan(&updatedTeam).Error; err != nil {
        log.Println(err.Error())
        return nil, errors.New("failed to update team")
    }
	
	newValues := updatedTeam
	//update audit log
	if err := utilities.LogAudit("Update",user_id,role,"Team",teamID,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
    return &existingTeam, nil
}
/*
deletes team
@params team_id
*/
func DeleteTeam(c *fiber.Ctx, team_id uuid.UUID) error {
	user_id , _ := GetAuthUserID(c)
	role := GetAuthUser(c)
    // Fetch the current state of the team before deletion
    team := new(Team)
    err := db.First(&team, "id = ?", team_id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Println(err.Error())
            return errors.New("record not found")
        }
        log.Println(err.Error())
        return errors.New("failed to fetch team")
    }
	oldValues := team

    // Delete the team record from the database
    err = db.Delete(&team).Error
    if err != nil {
        log.Println(err.Error())
        return errors.New("failed to delete team")
    }

    // Log the audit event for the delete operation
    err = utilities.LogAudit(
        "delete",
        user_id,   
        role,      
        "Team",        
        team_id,       
        oldValues,          
        nil,           
        c,             
    )
    if err != nil {
        log.Println("error logging audit:", err.Error())
    }

    return nil
}
