package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	//"github.com/jinzhu/gorm"
	"main.go/database"
	"main.go/model"
	"main.go/utilities"
)


var db =database.ConnectDB()
var country_code = "KE"

func CreateUserAccount(c *fiber.Ctx) error {
	//generating new id
	id := uuid.New()
	db.AutoMigrate(&model.User{})
	user:=model.User{}
	if err :=c.BodyParser(&user); err != nil {
		return utilities.ShowError(c,"failed to parse JSON data", fiber.StatusInternalServerError)
	}

	//validate email address
	if ! utilities.ValidateEmail(user.Email){
		return utilities.ShowError(c,"inavalid email address", fiber.StatusNotAcceptable)
	}
	//Check if user exist
	userExist,_,_ := model.UserExist(c,user.PhoneNumber)
	if userExist{
		return utilities.ShowError(c,"User with this phone number already exists",fiber.StatusConflict)
	}
	//validate phone number
	_,err := utilities.ValidatePhoneNumber(user.PhoneNumber,country_code)
	if err !=nil{
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
	result := db.Create(&userModel)
	if result.Error != nil {
		return utilities.ShowError(c, "failed to add data to the database",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"account created successfully",fiber.StatusOK)
}

