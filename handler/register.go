package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"main.go/database"
	"main.go/model"
	"main.go/utilities"
)


var db =database.ConnectDB()


func CreateUserAccount(c *fiber.Ctx) error {
	//generating new id
	id := uuid.New()
	
	user:=model.User{}
	if err :=c.BodyParser(&user); err != nil {
		return c.JSON(fiber.Map{"error":err.Error()})
	}


	email:=user.Email
	
	//validate email address
	if ! utilities.ValidateEmail(email){
		return utilities.ShowError(c,"inavalid email address", fiber.StatusNotAcceptable)
	}

	///check if the user is already registered
	result :=db.Where("phone_number = ?",user.PhoneNumber).First(&user)
	if result.Error == nil{
		return utilities.ShowError(c,"User with this phone number already exists",fiber.StatusConflict)
	}else if result.Error != gorm.ErrRecordNotFound{
		return utilities.ShowError(c,"internal server error",fiber.ErrNotFound.Code)
	}

	//validate phone number
	_,err := utilities.ValidatePhoneNumber(user.PhoneNumber,user.CountryCode)
	if err !=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusAccepted)
	}

	//comapare passwords
	if user.ConfirmPassword != user.Password{
		return utilities.ShowError(c,"passwords do not match",fiber.StatusForbidden)
	}
	
	//hash password
	hashed_password, _:= utilities.HashPassword(user.Password)

	userModel := model.User{
		ID: id,
		FullName: user.FullName,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		CountryCode: user.CountryCode,
		Password: hashed_password,
	}
	db.AutoMigrate(&userModel)
	//create user
	result = db.Create(&userModel)
	if result.Error != nil {
		return utilities.ShowError(c, "failed to add data to the database",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"account created successfully",fiber.StatusOK)
}

