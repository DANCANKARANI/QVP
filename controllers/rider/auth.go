package rider

import (
	"log"
	"time"

	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/DANCANKARANI/QVP/middleware"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

type loginResponse struct{
	Token string  `json:"token"`
}
//rider login handler
func RiderLogin(c *fiber.Ctx) error {
	rider := new(model.Rider)

	//parse request body
	if err := c.BodyParser(rider); err != nil{
		log.Println("error parsing request body",err.Error())
		return utilities.ShowError(c,"failed to log in",fiber.StatusInternalServerError)
	}

	//check if user exists
	exist,existingUser,_:=model.RiderExist(c,rider.PhoneNumber)
	if !exist{
		err_str :="user with phone number:"+rider.PhoneNumber+" doesn't exist"
		return utilities.ShowError(c,err_str,fiber.StatusInternalServerError)
	}


	//verify logins
	err :=utilities.CompareHashAndPassowrd(existingUser.Password,rider.Password)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}

	
	exp :=time.Hour*24
	//generating token
	tokenString,err := middleware.GenerateToken(middleware.Claims{UserID: &existingUser.ID,Role: "rider"},exp)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	//set token cookie 
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24), // Same duration as the token
		HTTPOnly: true, // Important for security, prevents JavaScript access
		Secure:   true, // Use secure cookies in production
		Path:     "/",  // Make the cookie available on all routes
	})
	 
	response_user:=loginResponse{
		Token: tokenString,
	}

	//update audit logs
	if err := utilities.LogAudit("Register",existingUser.ID,"Rider","Branch",existingUser.ID,nil,nil,c); err != nil{
		log.Println(err.Error())
	}

	//return
	return utilities.ShowSuccess(c,"successfully logged in",fiber.StatusOK,response_user)
}

//rider logout handler
func Logout(c *fiber.Ctx) error {
	
	//call service logout
	err := user.LogoutService(c,"Rider")
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	//response
	return utilities.ShowMessage(c,"successfully logged out",fiber.StatusOK)
}
