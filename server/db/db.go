package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {

	var err error
	db, err = gorm.Open("mysql", "root@/daigo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	createTable()
}

func createTable() {
	if !db.HasTable(&Product{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Product{}).Error; err != nil {
			panic(err)
		}
	}
}
