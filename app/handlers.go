package app

import (
	"encoding/json"
	"encoding/xml"
	"go-banking/service"
	"net/http"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

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
