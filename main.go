package main

import (
	"github.com/DANCANKARANI/QVP/endpoints"
	"github.com/DANCANKARANI/QVP/model"
)
func main() {
	model.Migration()
	endpoints.CreateEndpoint()
}