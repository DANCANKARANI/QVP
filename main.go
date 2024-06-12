package main

import (

	"main.go/routes"
	"main.go/database"
)


func main() {
	
	database.ConnectDB()
	routes.EndPoints()
}