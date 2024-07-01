package main

import(
	"main.go/endpoints"
	"main.go/model"
)

func main() {
	endpoints.CreateEndpoint()
	model.Migration()
	//pkg.Start()
}