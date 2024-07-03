package main

import(
	"github.com/DANCANKARANI/QVP/endpoints"
	//"github.com/DANCANKARANI/QVP/model"
	"github.com/DANCANKARANI/QVP/pkg"
)

func main() {
	//model.Migration()
	pkg.Find()
	endpoints.CreateEndpoint()

	//pkg.Start()
}