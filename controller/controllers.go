package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lanpaiva/person-api-rest/models"
)

func HandlerRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
	fmt.Println(w, "iniciando o servidor Rest com Go")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Persons)
}

func IdPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, person := range models.Persons {
		if strconv.Itoa(person.ID) == id {
			json.NewEncoder(w).Encode(person)
		}
	}
}
