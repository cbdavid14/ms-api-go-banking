package app

import (
	"encoding/json"
	"github.com/cbdavid14/ms-api-go-banking/dto"
	"github.com/cbdavid14/ms-api-go-banking/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	service service.AccountService
}

func (ch AccountHandler) save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var req dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.CustomerId, _ = strconv.Atoi(customerId)
		response, appError := ch.service.Save(req)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}

func (ch AccountHandler) makeTransaction(w http.ResponseWriter, r *http.Request) {
	//get params from path
	vars := mux.Vars(r)
	customerId, _ := strconv.Atoi(vars["customer_id"])
	accountId, _ := strconv.Atoi(vars["account_id"])

	//decode incoming request
	var req dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.CustomerId = customerId
		req.AccountId = accountId
		response, appError := ch.service.MakeTransaction(req)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}

}
