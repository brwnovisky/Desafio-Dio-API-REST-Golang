package main

import (
	"Desafio-Dio-API-REST-Golang/context"
	"Desafio-Dio-API-REST-Golang/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	context := context.NewDatabaseContext("clients")

	homeController :=
		controllers.RequestHandler{IController: controllers.NewHome()}
	clientController :=
		controllers.RequestHandler{IController: controllers.NewClient(context)}
	clientIdController :=
		controllers.RequestHandler{IController: controllers.NewClientId(context)}

	router := mux.NewRouter()
	router.HandleFunc("/", homeController.RequestHandler)
	router.HandleFunc("/client", clientController.RequestHandler)
	router.HandleFunc("/client/{id}", clientIdController.RequestHandler)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
