package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lanpaiva/person-api-rest/models"
)

func HandlerRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
	fmt.Println(w, "iniciando o servidor Rest com Go")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Persons)
}
