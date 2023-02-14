package routes

import (
	"log"
	"net/http"

	"github.com/lanpaiva/person-api-rest/controller"
)

func HandleResponse() {
	http.HandleFunc("/", controller.HandlerRequest)
	http.HandleFunc("/persons", controller.AllPersons)
	defer log.Fatal(http.ListenAndServe(":8080", nil))
}
