package config

import (
	"github.com/jinzhu/gorm" //Mengkases Database
	_ "github.com/jinzhu/gorm/dialects/mysql" //Mengimpor driver MySQL 
)

var (
	db * gorm.DB
)

func connect(){
	d, err := gorm.Open("mysql", ":?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}