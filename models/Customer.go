package models

import (
	exp "Exception"
	database "SqlOperations"
	"github.com/jinzhu/gorm"
	"regexp"
)

var db *gorm.DB

type Customer struct {
	//Key97    int    `gorm:"primary_key;not null;AUTO_INCREMENT"json:"key97"`
	Mnemonic string `gorm:"primary_key;not null;index:CUST_KEY"json:"mnemonic"`
	Name     string "json:name"
	Address  string "json:address"
	Phone    string "json:phone"
}

func init() {
	database.DBConnect()
	db = database.GetDBConnection()
	db.AutoMigrate(&Customer{})
	//db.CreateTable(&Customer{})
}

func GetAllCustomers() []Customer {
	var Cust []Customer
	db.Find(&Cust)
	return Cust
}

func GetCustomerById(CustId string) (*Customer, *gorm.DB) {
	var getCust Customer
	db := db.Find(&getCust, "mnemonic = ?", CustId)
	return &getCust, db
}

func (e *Customer) CreateCustomer() (*Customer, *exp.ErrorHandler) {

	boolVal, errorVal := e.ValidateCustomerDetails()
	if boolVal == false {
		errorVal = nil
		db.NewRecord(e)
		db.Create(e)
	}

	return e, errorVal
}

func DeleteCustomer(CustId string) Customer {
	var cust Customer
	db.Where("mnemonic = ?", CustId).Delete(cust)
	return cust
}

func (cust *Customer) ValidateCustomerDetails() (bool, *exp.ErrorHandler) {
	//Phone Number Validation
	match, _ := regexp.MatchString("([a-z]+)", cust.Phone)
	if match {
		err := exp.SetErrorDetails("ERR_003", "Phone Number can be only of Numbers")
		return true, err
	}
	if len(cust.Phone) != 10 {
		err := exp.SetErrorDetails("ERR_002", "Phone number cannot be Greator/Less than 10 Digit")
		return true, err
	}

	//Customer Validation
	if len(cust.Mnemonic) > 20 {
		err := exp.SetErrorDetails("ERR_001", "Customer Mnemonic cannot be greator than 20 Charctaers")
		return true, err
	}

	return false, nil

}
