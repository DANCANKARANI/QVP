package main

import (
	"main.go/database"
	"main.go/routes"
)


func main() {
	database.RedisClient()
	database.ConnectDB()
	routes.EndPoints()
}