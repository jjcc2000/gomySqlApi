package db

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Using the gorm module to conect to mySQL
//go get -u gorm.io/driver/mysql for mySql
var stnDB ="root:sqlpassword@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB


func DBConection() {
	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	var err error
	DB , err = gorm.Open(mysql.Open(stnDB),&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}else{
			log.Println("Is conected succesfully")
	}

}
