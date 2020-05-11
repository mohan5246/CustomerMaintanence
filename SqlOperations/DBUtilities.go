package SqlOperations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func DBConnect() {
	// Please define your user name and password for my sql.
	d, err := gorm.Open("mysql", "root@/employeedb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDBConnection() *gorm.DB {
	return db
}
