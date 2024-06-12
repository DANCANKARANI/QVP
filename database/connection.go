package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
    fmt.Println(err)
	}

	user_name:=os.Getenv("DB_USER")
	password:= os.Getenv("DB_PASSWORD")
	dsn := user_name+":"+password+ "@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
		
	}
	fmt.Println("Connected...")
	return db
}