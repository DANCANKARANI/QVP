package user

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	//"github.com/jinzhu/gorm"
	"github.com/DANCANKARANI/QVP/database"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
)


var db =database.ConnectDB()
var country_code = "KE"

func CreateUserAccount(c *fiber.Ctx) error {
	db.AutoMigrate(&model.User{})
	//generating new id
	id := uuid.New()
	user:=model.User{}
	if err :=c.BodyParser(&user); err != nil {
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to create account", fiber.StatusInternalServerError)
	}

	//validate email address
	_,err:=utilities.ValidateEmail(user.Email)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	//Check if user exist
	userExist,_,err:= model.UserExist(c,user.PhoneNumber)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	if userExist{
		return utilities.ShowError(c,"User with this phone number already exists"+user.PhoneNumber,fiber.StatusConflict)
	}
	//validate phone number
	phone,err := utilities.ValidatePhoneNumber(user.PhoneNumber,country_code)
	if err !=nil || phone ==""{
		log.Println(err.Error())
		return utilities.ShowError(c,err.Error(),fiber.StatusAccepted)
	}

	//comapare passwords
	if user.ConfirmPassword != user.Password{
		return utilities.ShowError(c,"passwords do not match",fiber.StatusForbidden)
	}
	
	//hash password
	hashed_password, _:= utilities.HashPassword(user.Password)

	userModel := model.User{ID: id,FullName: user.FullName,Email: user.Email,PhoneNumber: user.PhoneNumber,CountryCode: country_code,Password: hashed_password,ResetCode: "",}
	//create user
	userModel.CodeExpirationTime=time.Now()
	err = db.Create(&userModel).Error
	if err!= nil {
		log.Fatal(err.Error())
		return utilities.ShowError(c, "failed to add data to the database",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"account created successfully",fiber.StatusOK)
}

