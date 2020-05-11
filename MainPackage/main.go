package main

import (
	"Handler"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	
	r := mux.NewRouter()
	Handler.CustomerMaintanence(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8005", r))
}
