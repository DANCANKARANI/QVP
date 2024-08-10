package model

import (
	"errors"
	"log"

	"github.com/DANCANKARANI/QVP/utilities"
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

// UpdateUser updates the user by ID and logs the changes.
func UpdateUser(c *fiber.Ctx) (*ResponseUser, error) {
    // Get the authenticated user ID
	role :=GetAuthUser(c)
    id, err := GetAuthUserID(c)
    if err != nil {
        return nil, errors.New("failed to get user's ID: " + err.Error())
    }

    // Parse the request body into a User struct
    var body User
    if err := c.BodyParser(&body); err != nil {
        return nil, errors.New("failed to parse: " + err.Error())
    }

    // Fetch the current user record to get old values
    oldValues := new(User)
    if err := db.First(&oldValues, "id = ?", id).Error; err != nil {
        return nil, errors.New("failed to fetch current user: " + err.Error())
    }
	response := new(ResponseUser)
    // Update the user record
    if err := db.Model(&User{}).Where("id = ?", id).Updates(&body).Scan(&body).Scan(response).Error; err != nil {
        return nil, errors.New("error in updating the user: " + err.Error())
    }
    // Audit logs
    newValues := &body
    if err := utilities.LogAudit("update", id, role, "User", id, oldValues, newValues, c); err != nil {
        log.Println(err.Error())
        return nil, errors.New("error updating audit log")
    }

    return response, nil
}
/*
adds users insurance 
@params user_id, insurance_id
*/

func AddUserInsurance(c *fiber.Ctx,user_id,insurance_id uuid.UUID)(*InsuranceUser,error){
	role :=GetAuthUser(c)
	id:=uuid.New()
	insuranceUser:=InsuranceUser{
		ID:id,
		UserID: user_id,
		InsuranceID:insurance_id,
	}
	oldValues := ""
	err:=db.Create(&insuranceUser).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to add insurance users")
	}
	newValues :=insuranceUser
	if err = utilities.LogAudit("Create",user_id,role,"Insurance",insurance_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return &insuranceUser,nil
}

/*
updates user insurance
@params user_id, insurance_id
*/
func UpdateUserInsurance(c *fiber.Ctx,id,user_id,insurance_id uuid.UUID) (*InsuranceUser,error){
	role :=GetAuthUser(c)
	insuranceUser := InsuranceUser{
		UserID: user_id,
		InsuranceID: insurance_id,
	}
	oldValues := new(InsuranceUser)
	err:=db.First(&oldValues,"id = ?",id).Updates(&insuranceUser).Scan(&insuranceUser).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update insurance users")
	}
	newValues := insuranceUser
	if err = utilities.LogAudit("Update",user_id,role,"Insurance",insurance_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
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
