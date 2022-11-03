package app

import (
	"encoding/json"
	"go-banking/dto"
	"go-banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = customerId
	account, appErr := h.service.NewAccount(request)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.Message)
		return
	}

	writeResponse(w, http.StatusCreated, account)
}