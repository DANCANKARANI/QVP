package main

import (
	"main.go/database"
	"main.go/model"
	"main.go/routes"
)


func main() {
	database.ConnectDB()
	model.Migration()
	routes.EndPoints()
}