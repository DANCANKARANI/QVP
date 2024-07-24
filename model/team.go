package model

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//creates a new team
func CreateTeam(c *fiber.Ctx) (*Team, error) {
 team := new(Team)
 if err := c.BodyParser(&team); err != nil {
	log.Println(err.Error())
	return nil, errors.New("failed to parse json data")
 }
 team.ID = uuid.New()
 err := db.Create(&team).Error
 if err != nil {
	return nil, errors.New("failed to create team")
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
func UpdateTeam(c *fiber.Ctx,team_id uuid.UUID)(*Team,error){
	team := new(Team)
	if err:=c.BodyParser(&team); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to parse json data")
	}
	if err := db.First("id = ?", team_id).Updates(&team).Scan(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return nil,errors.New("failed to parse json data")
		}
		log.Println(err.Error())
		return nil,errors.New("failed to update team")
	}
	return team,nil
}
/*
deletes team
@params team_id
*/
func DeleteTeam(team_id uuid.UUID)error{
	team := new(Team)
	err :=db.First(&team).Delete(&team).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			log.Println(err.Error())
			return errors.New("record not found")
		}
		log.Println(err.Error())
		return errors.New("failed to delete team")
	}
	return nil
}