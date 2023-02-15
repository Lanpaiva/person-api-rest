package models

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Historia string `json:"historia"`
}

var Persons []Person
