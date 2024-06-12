package middleware

import (
	"fmt"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

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