package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"main.go/database"
	"main.go/model"
	"main.go/utils"
)


var db =database.ConnectDB()


func CreateUserAccount(c *fiber.Ctx) error {
	//generating new id
	id := uuid.New()
	
	user:=model.User{}
	if err :=c.BodyParser(&user); err != nil {
		return c.JSON(fiber.Map{"error":err.Error()})
	}

	//check if the user is already registered
	email:=user.Email
	
	//validate email address
	if ! utils.ValidateEmail(email){
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message":"invalid email address"})
		}else{
			result :=db.Where("phone_number = ?",user.PhoneNumber).First(&user)
				if result.Error == nil{
					//email already exist in the database
					return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message":"User with this phone number already exists"})
				}else if result.Error != gorm.ErrRecordNotFound{
					c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"message":"internal server error"})
				}
			_,err := utils.ValidatePhoneNumber(user.PhoneNumber,user.CountryCode)
			if err !=nil{
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message":err.Error()})
			}else{
				if user.ConfirmPassword != user.Password{
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message":"passwords do not match",
					})
				}else{
					hashed_password, _:= utils.HashPassword(user.Password)

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
						return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
							"error":result.Error,
						})
					}
					return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"user created successfully"})
				}
		
			}
		}
	

}

