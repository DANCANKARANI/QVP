package main

import(
	"github.com/DANCANKARANI/QVP/endpoints"
	"github.com/DANCANKARANI/QVP/model"
	//"github.com/DANCANKARANI/QVP/database"
)
func main() {
	model.Migration()
	model.InitFirebase()
	//database.StartRedisServer()
	endpoints.CreateEndpoint()
}