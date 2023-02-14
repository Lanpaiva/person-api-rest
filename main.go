package main

import (
	"github.com/lanpaiva/person-api-rest/models"
	"github.com/lanpaiva/person-api-rest/routes"
)

func main() {
	models.Persons = []models.Person{
		{Name: "Nome 1", Historia: "Historia 1"},
		{Name: "Nome 2", Historia: "Historia 2"},
	}

	routes.HandleResponse()
}
