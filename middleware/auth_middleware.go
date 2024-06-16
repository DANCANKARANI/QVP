package middleware

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	//"main.go/database"
	"main.go/model"
	//"main.go/database"
)

type Claims struct {
	UserID *uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(c *fiber.Ctx, id string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	my_secret_key := os.Getenv("MY_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":       id,
		"expiration_at": time.Now().Add(time.Minute * 60).Unix(), //expires after an hour
	})
	tokenString, err := token.SignedString([]byte(my_secret_key))
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func GenerateToken(claims Claims, exp time.Duration) (string, error) {


	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(exp).Unix(),
		Issuer:    "qvp",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("MY_SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//validate token 
func ValidateToken(tokenString string)(*Claims,error){
	token,err := jwt.ParseWithClaims(tokenString, &Claims{},func(token *jwt.Token) (interface{}, error) {
		return []byte("MY_SECRET_KEY"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims,ok := token.Claims.(*Claims); ok && token.Valid{
		isRevoked :=model.IsRevoked(tokenString)
		if !isRevoked {
			
			return nil, errors.New("user token is revoked")
		}
		return claims, nil
	}
	return nil,errors.New("invalid token")
}

// func InvalidateToken(tokenString string) error {
// 	if err :=db.AutoMigrate(&model.RevokedToken{});err != nil{
// 		return err.Error
// 	}
// 	result :=db.Create(&model.RevokedToken{
// 		Token: tokenString,
// 		RevokedAt: time.Now(),
// 	})
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	fmt.Println(result)
// 	return nil
// }

func GetAuthUserID(c *fiber.Ctx,claims *Claims)(*uuid.UUID,error){
	if claims == nil{
		return nil, errors.New("unauthorized user denied. user details not found")
	}
	//extract the user ID from the claims
	userClaimID := claims.UserID
	if userClaimID ==nil{
		return nil, errors.New("unauthorized user denied. user details not found")
	}
	return userClaimID, nil
}