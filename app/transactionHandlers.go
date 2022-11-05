package app

import (
	"encoding/json"
	"go-banking/dto"
	"go-banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionHandlers struct {
	service service.TransactionService
}

func (th TransactionHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	account_id := vars["account_id"]

	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.AccountId = account_id
	request.CustomerId = customer_id

	t, appErr := th.service.MakeTransaction(request)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w,http.StatusOK, t)
}
