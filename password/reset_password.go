package password

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/model"
	"main.go/utilities"
)
var db = database.ConnectDB()
func ResetPassword(c *fiber.Ctx,email, phone_number,password,code string) {
	user := model.User{}
	
	user, _ = model.FindUser(email, phone_number)
	if code != user.ResetCode || time.Now().After(user.CodeExpirationTime) {
		utilities.ShowError(c, "invalid reset code, request another code", fiber.StatusNotAcceptable)
	}
	user.Password, _ = utilities.HashPassword(password)
	fmt.Println(user.Password)
	db.Save(&user)
}
