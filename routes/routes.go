package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanpaiva/person-api-rest/controller"
)

func HandleResponse() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.HandlerRequest)
	r.HandleFunc("/persons", controller.AllPersons)
	defer log.Fatal(http.ListenAndServe(":8080", r))
}
