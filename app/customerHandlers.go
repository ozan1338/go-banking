package app

import (
	"encoding/json"
	"encoding/xml"
	"go-banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Aish", City: "Japan", Zipcode: "110075"},
	// 	{Name: "Rob", City: "Korea", Zipcode: "11076"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer,err := ch.service.GetCustomer(id)

	if err != nil {
		// w.Header().Add("Content-type", "application/json")
		// w.WriteHeader(err.Code)
		// fmt.Fprint(w, err.Message)
		// json.NewEncoder(w).Encode(err.AsMessage())
		message := err.AsMessage()
		writeResponse(w, err.Code, message)
		return
	}

	// w.Header().Add("Content-type", "application/json")
	// json.NewEncoder(w).Encode(customer)
	writeResponse(w,http.StatusOK,customer)
}

func writeResponse(w http.ResponseWriter,code int, data any ) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
