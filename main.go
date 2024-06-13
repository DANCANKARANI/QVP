package main

import (
	"main.go/database"
	"main.go/routes"
)


func main() {
	
	database.ConnectDB()
	routes.EndPoints()
}