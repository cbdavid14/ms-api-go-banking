package app

import (
	"encoding/json"
	"github.com/cbdavid14/ms-api-go-banking/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	var customers, err = ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch CustomerHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId, _ := strconv.Atoi(vars["customer_id"])

	var customer, err = ch.service.GetCustomerById(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
