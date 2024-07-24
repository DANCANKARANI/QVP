
package notification

import (

	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/utilities"

)

	func AddNotification(c *fiber.Ctx) error{
		body:=model.Notification{}
		if err := c.BodyParser(&body); err != nil {
			return utilities.ShowError(c,"failed to parse json data", fiber.StatusInternalServerError)
		}
		notification,err := model.SendNotification(c,body)
		if err != nil {
			return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
		}
		return utilities.ShowSuccess(c,"successfully added notification",fiber.StatusOK,notification)
	}

	func GetNotification(c *fiber.Ctx)error{
		id,_:=model.GetAuthUserID(c)
		notification, err := model.GetNotification(c,id)
		if err != nil {
			return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
		}
		return utilities.ShowSuccess(c,"successfully retrieved notification",fiber.StatusOK,notification)
	}

/*package notification

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)
  
 func InitializeNotifications(){
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil{
		fmt.Println("error innitializing the app: ",err.Error())
	}
	//Get fgm client
	client, err := app.Messaging(context.Background())
	if err != nil{
		fmt.Println("error getting fgm client: ",err.Error())
	}

	//send a message
	message :=  &messaging.Message{
		Data : map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		Notification: &messaging.Notification{
			Title: "Hello from Go",
			Body:  "This is a Firebase Cloud Messaging (FCM) notification message!",
		},
		Token: "<FCM registration token of the device>",
	}
	// Send message
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("error sending message: %v\n", err)
	}

	// Response is a message ID string
	log.Printf("Successfully sent message: %s\n", response)
 }
  
*/