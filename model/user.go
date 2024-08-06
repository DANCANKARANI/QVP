package model

import (
	"errors"
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
gets user using user ID
*/
type ResponseUser struct{
	ID uuid.UUID 		`json:"id"`
	FullName string 	`json:"full_name"`
	PhoneNumber string 	`json:"phone_number"`
	Email string 		`json:"email"`
	ImageID uuid.UUID   `json:"image_id,omitempty"`
}

func GetOneUSer(c *fiber.Ctx)(*ResponseUser,error){
	id,err:=GetAuthUserID(c)
	if err != nil {
		return nil,errors.New("failed to get user's id:"+err.Error())
	}
	user := ResponseUser{}
	err = db.Preload("Image").First(&User{},"id = ?",id).Scan(&user).Error
	if err != nil {
		return nil,errors.New("failed to get user details:"+err.Error())
	}
	return &user,nil
}
//gets all the users
func GetAllUsersDetails(c *fiber.Ctx)(*[]ResponseUser,error){
	response:=[]ResponseUser{}
	err := db.Model(&User{}).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to get users:"+err.Error())
	}
	return &response,nil
}

//updates the user by id
func UpdateUser(c *fiber.Ctx)(*ResponseUser,error){
	id,err:=GetAuthUserID(c)
	if err != nil {
		return nil,errors.New("failed to get user's id:"+err.Error())
	}
	body := User{}
	if err = c.BodyParser(&body);err != nil {
		return nil,errors.New("failed to parse:"+err.Error())
	}
	response := ResponseUser{}
	err = db.First(&User{},"id = ?",id).Updates(&body).Scan(&response).Error
	if err != nil {
		return nil,errors.New("error in updating the user:"+err.Error())
	}
	return &response,nil
}



func AddUserInsurance(user_id,insurance_id uuid.UUID)(*InsuranceUser,error){
	id:=uuid.New()
	insuranceUser:=InsuranceUser{
		ID:id,
		UserID: user_id,
		InsuranceID:insurance_id,
	}

	err:=db.Create(&insuranceUser).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to add insurance users")
	}
	return &insuranceUser,nil
}
func UpdateUserInsurance(id,user_id,insurance_id uuid.UUID) (*InsuranceUser,error){
	insuranceUser := InsuranceUser{
		UserID: user_id,
		InsuranceID: insurance_id,
	}
	fmt.Println(insurance_id)
	err:=db.First(&InsuranceUser{},"id = ?",id).Updates(&insuranceUser).Scan(&insuranceUser).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update insurance users")
	}
	return &insuranceUser,nil
}

//
func MapUserToResponse(user User) ResponseUser {
    return ResponseUser{
        ID:          user.ID,
        FullName:    user.FullName,
        PhoneNumber: user.PhoneNumber,
        Email:       user.Email,
        ImageID:     *user.ImageID, // Ensure this is populated correctly
    }
}
