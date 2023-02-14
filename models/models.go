package models

type Person struct {
	Name     string `json:"name"`
	Historia string `json:"historia"`
}

var Persons []Person
