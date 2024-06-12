package handler

import (

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"main.go/model"
	"main.go/utilities"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)


func Login(c *fiber.Ctx)error{
	user := model.User{}
	if err := c.BodyParser(&user); err !=nil {
		return c.JSON(fiber.Map{"error":err.Error()})
	}
	//check if the user exist
	var existingUser model.User
	result := db.Where("phone_number = ?",user.PhoneNumber).First(&existingUser)
	if result.Error != nil {
		//user not found
		return c.JSON(fiber.Map{"error":"invalid credintials 1"})
	}else{
		//compare password
		err :=utils.CompareHashAndPassowrd(existingUser.Password,user.Password)
		if err !=nil{
			return c.JSON(fiber.Map{"error":"invalid credintials 2"})
		}else{
			tokenString,err := GenerateJWT(c,existingUser.ID.String())
			if err != nil{
				return c.JSON(fiber.Map{"error":"failed to generate token"})
			}else{
				c.Set("Authorization","Bearer"+tokenString)
				return c.JSON(fiber.Map{"message":"successfully logged in"})
			}
		}
	}
}

func GenerateJWT(c *fiber.Ctx,id string)(string,error){
	err := godotenv.Load(".env")
	if err != nil {
    fmt.Println(err)
	}
	my_secret_key := os.Getenv("MY_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":"id",
		"exp":time.Now().Add(time.Minute*60).Unix(),//expires after an hour
	})
	tokenString, err := token.SignedString([]byte(my_secret_key))
	if err != nil {
		return "", err
	}else{
		return tokenString, nil
	}
}