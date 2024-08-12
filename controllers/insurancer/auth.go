package insurancer

import (
	"log"
	"time"

	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/DANCANKARANI/QVP/middleware"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/gofiber/fiber/v2"
)

//response
type loginResponse struct{
	Token string 	`json:"token"`
}
//insurancer login
func InsurancerLogin(c *fiber.Ctx) error {
	user := model.Insurancer{} 

	//parse request body
	if err := c.BodyParser(&user); err !=nil {
		return utilities.ShowError(c,"failed to login",fiber.StatusInternalServerError)
	}

	//check of user exist
	userExist,existingUser,_:= model.InsurerExist(c,user.PhoneNumber)
	if ! userExist {
		err_str :="insurancer with this phone number:"+ user.PhoneNumber+" does not exist"
		return utilities.ShowError(c,err_str,fiber.StatusNotFound)
	}
	
	//compare password
	err :=utilities.CompareHashAndPassowrd(existingUser.Password,user.Password)
	if err !=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusForbidden)
			 
	}
	exp :=time.Hour*24
	//generating token
	tokenString,err := middleware.GenerateToken(middleware.Claims{UserID: &existingUser.ID,Role: "Insurancer"},exp)
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
	if err := utilities.LogAudit("Login",existingUser.ID,"Insurancer","User",existingUser.ID,existingUser,existingUser,c); err != nil{
		log.Println(err.Error())
	}

	return utilities.ShowSuccess(c,"successfully logged in",fiber.StatusOK,response_user)
}

//logout
func Logout(c *fiber.Ctx) error {
	//call logout service
	err := user.LogoutService(c,"Insurancer")
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	//response
	return utilities.ShowMessage(c,"successfully logged out",fiber.StatusOK)
}
