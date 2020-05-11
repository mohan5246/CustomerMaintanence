package Routers

import (
	"MessageParser"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"models"
	"net/http"
)

var NewCust models.Customer

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	CreateCustomer := &models.Customer{}
	MessageParser.ParseBody(r, CreateCustomer)
	b, erx := CreateCustomer.CreateCustomer()
	if erx != nil {
		res, _ := json.Marshal(erx)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		res, _ := json.Marshal(b)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	cust := models.GetAllCustomers()
	res, _ := json.Marshal(cust)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CustId := vars["custId"]
	log.Println("Test", CustId)
	customerDetails, _ := models.GetCustomerById(CustId)
	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCustomerDetails(w http.ResponseWriter, r *http.Request) {
	var updateDetails = &models.Customer{}
	MessageParser.ParseBody(r, updateDetails)
	vars := mux.Vars(r)
	custId := vars["custId"]

	customerDetails, db := models.GetCustomerById(custId)
	if updateDetails.Name != "" {
		customerDetails.Name = updateDetails.Name
	}
	if updateDetails.Address != "" {
		customerDetails.Address = updateDetails.Address
	}
	if updateDetails.Phone != "" {
		customerDetails.Phone = updateDetails.Phone
	}
	db.Save(&customerDetails)
	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custId := vars["custId"]
	
	cust := models.DeleteCustomer(custId)
	res, _ := json.Marshal(cust)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
