package pkg

import (
	"fmt"

	//"github.com/DANCANKARANI/QVP/controllers/image"
	"github.com/DANCANKARANI/QVP/database"
	"github.com/DANCANKARANI/QVP/model"
)

type Shape interface {
	GetArea() float64
	GetDescription(dimension string) error
}

type TriangleShape struct {
	Height float64
	Base   float64
}

func (t TriangleShape) GetArea() float64 {
	return 0.5 * t.Height * t.Base
}
func (t TriangleShape) GetDescription(dimension string) error {
	_,err:=fmt.Printf("Triangle has a base of %v and height of %v", t.Base,t.Height)
	return err
}

type RectangleShape struct {
	Length float64
	Width float64
}
func (R RectangleShape) GetArea() float64 {
	return R.Length*R.Width
}
func (R RectangleShape) GetDescription(dimension string) error {
	fmt.Printf(" %s: dimension %v\n",  dimension,R.Length)
	 return nil
}
func SendShape(s Shape,dimension string) error {
	return s.GetDescription(dimension)
}

//interface
type Notifier interface {
	Notify(id string,message interface{}) error
}

type InApp struct{
	ID string
}
func (in InApp) Notify(message interface{}) error {
	fmt.Printf("Message: %s for user with id %s:",message,in.ID)
	return nil
}
type SMS struct{
	Phone string
}
func (s SMS) Notify(id string,message interface{})error{
	fmt.Printf("sending %s to this phone number %s",message,s.Phone)
	return nil
}

func SendNotification(n Notifier,id,message string) error{
	return n.Notify(id,message)
}
var number ="0797408042"
func SendSms(){
	sms := SMS{Phone: number}
	SendNotification(sms,number,number)
}





var db = database.ConnectDB()
type Notification interface{
	Notify(id string)(*[]model.Notification,error)
}

type NormalUser struct {}




func Start(){
	// rectangle:=RectangleShape{Width:20,Length:50}
	// triangle := TriangleShape{Base: 10,Height:20}
	// SendShape(triangle,"start")
	// SendShape(rectangle,"finish")
	//inApp := InApp{ID:"1"}

}


func Run(){
	db:=database.ConnectDB()
	db.AutoMigrate(&model.Admin{},&model.User{})
}
func Find(){
	
	user:=model.User{}
	err := db.Where("phone_number = ?", "0797409042").First(&user).Error
	if err != nil {
    fmt.Println(err.Error())
	}
	fmt.Println(user)
}