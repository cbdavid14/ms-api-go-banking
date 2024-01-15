package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	//define a route
	//mux := http.NewServerMux()
	router := mux.NewRouter()

	//define the handlers
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", createUser).Methods(http.MethodPost)

	router.HandleFunc("/users/{id:[0-9]+}", getUserById).Methods(http.MethodGet)
	//start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Post request received")
}

func getUserById(writer http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(writer, vars["id"])
}
