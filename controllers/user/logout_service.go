package user

import (
	"log"
	"time"

	"github.com/DANCANKARANI/QVP/middleware"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

func LogoutService(c *fiber.Ctx, user_type string) error {
	user_id, _ := model.GetAuthUserID(c)
	role := model.GetAuthUser(c)

	//get token string
	tokenString, err := utilities.GetJWTToken(c)
	if err != nil {
		return utilities.ShowError(c, err.Error(), fiber.StatusUnauthorized)
	}
	oldValues := tokenString

	//invalidate token
	err = middleware.InvalidateToken(tokenString)
	if err != nil {
		return utilities.ShowError(c, "failed to invalidate the token", fiber.StatusInternalServerError)
	}
	newValues := tokenString

	//update audit logs
	if err := utilities.LogAudit("Logout", user_id, role, user_type, user_id, oldValues, newValues, c); err != nil {
		log.Println(err.Error())
	}
	//set token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
	})

	//response
	return nil
}
