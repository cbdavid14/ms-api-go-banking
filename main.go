package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   int    `json:"id_user" xml:"id_xml"`
	Name string `json:"name_user" xml:"name_xml"`
}

func main() {
	//define a route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/users", getAllUsers)
	//start the server
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Id: 1, Name: "John Doe"},
		{Id: 2, Name: "Jane Doe"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(users)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
