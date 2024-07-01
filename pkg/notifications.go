package pkg

import (
	"errors"
	"fmt"

	"main.go/database"
	"main.go/model"
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
var notification  []model.Notification
type Notification interface{
	Notify(id string)(*[]model.Notification,error)
}

type NormalUser struct {}
func (u NormalUser)Notify(id string)(*[]model.Notification,error){
	exist := db.Find(&notification).Where("user_id = ?", "id").Scan(&notification).RecordNotFound()
	if ! exist{
		fmt.Println("failed", "Record not found")
		return nil,errors.New("record not found")
	}
	fmt.Println(notification)
	return &notification,nil
}

func sendNotification(N Notification,id string)(*[]model.Notification,error){
	return N.Notify(id)
}
func Send(){
	id := "533efd1a-5470-4579-87c8-f48c284e03b3"
	user:=NormalUser{}
	sendNotification(&user,id)	
}
func Start(){
	// rectangle:=RectangleShape{Width:20,Length:50}
	// triangle := TriangleShape{Base: 10,Height:20}
	// SendShape(triangle,"start")
	// SendShape(rectangle,"finish")
	//inApp := InApp{ID:"1"}
	Send()
}

