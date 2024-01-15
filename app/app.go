package app

import (
	"github.com/cbdavid14/ms-api-go-banking/domain"
	"github.com/cbdavid14/ms-api-go-banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//wiring instances
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
