package main

import (
	"api-go-rest/database"
	"api-go-rest/model"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	model.Personalities = []model.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}
	database.Connection()
	fmt.Println("Starting a rest server")
	routes.HandleRequest()
}
