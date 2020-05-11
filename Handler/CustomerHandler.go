package Handler

import (
	"Routers"
	"github.com/gorilla/mux"
)

var CustomerMaintanence = func(router *mux.Router) {
	router.HandleFunc("/customer/newcustomer/", Routers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/allcustomers/", Routers.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{custId}", Routers.GetCustomerById).Methods("GET")
	router.HandleFunc("/customer/updatedetails/{custId}", Routers.UpdateCustomerDetails).Methods("PUT")
	router.HandleFunc("/customer/deletecustomer/{custId}", Routers.DeleteCustomer).Methods("DELETE")
}
