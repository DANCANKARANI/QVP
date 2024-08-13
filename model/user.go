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
	ProfilePhotoPath string	`json:"profile_photo_path"`
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

	//validate phone number
	if body.PhoneNumber !=""{
		_,err :=utilities.ValidatePhoneNumber(body.PhoneNumber,country_code)
		if err != nil{
			return nil, err
		}
		exist,_,err:=UserExist(c,body.PhoneNumber)
		if err != nil{
			return nil, err
		}
		if exist{
			err_str := "user with phone:"+body.PhoneNumber+" already exist"
			return nil, errors.New(err_str)
		}
	}

	//validate email
	if body.Email !=""{
		_, err := utilities.ValidateEmail(body.Email)
		if err != nil{
			return nil, err
		}
	}

	//hash password
	if body.Password != ""{
		hashed_password, err:= utilities.HashPassword(body.Password)
		if err != nil{
			return nil,err
		}
		body.Password= hashed_password
	}
    // Fetch the current user record to get old values
    oldValues := new(User)
    if err := db.First(&oldValues, "id = ?", id).Error; err != nil {
        return nil, errors.New("failed to fetch current user: " + err.Error())
    }
	response := new(ResponseUser)

    // Update the user record
    if err := db.Model(&oldValues).Updates(&body).Scan(response).Error; err != nil {
        return nil, errors.New("error in updating the user: " + err.Error())
    }


    // Audit logs
    newValues := &oldValues
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
        ProfilePhotoPath: user.ProfilePhotoPath,
    }
}

/*
update users profile picture
@params user_id
*/
func UpdateUserProfilePic(c *fiber.Ctx, user_id uuid.UUID)(*ResponseUser,error){
	user := new(User)

	//generate image url
	profile_photo_path,err:=utilities.GenerateUrl(c,"profile")
	if err != nil{
		return nil, err
	}
	User := User{
		ProfilePhotoPath: profile_photo_path,
	}
	
	//find user
	if err := db.First(&user,"id = ?",user_id).Error; err != nil{
		log.Println("user not found:",err.Error())
		return nil, errors.New("failed to update profile image")
	}
	oldValues := user
	response := new(ResponseUser)
	//update profile image
	if err := db.Model(user).Updates(&User).Scan(response).Error; err != nil{
		log.Println("failed to update profile image:", err.Error())
		return nil, errors.New("failed to update profile image")
	}
	newValues := user
	
	//update audit log

	role := GetAuthUser(c)

	if err := utilities.LogAudit("Update",user_id,role,"User",user_id,oldValues,newValues,c); err != nil{
		log.Println(err.Error())
	}
	return response, nil
}
