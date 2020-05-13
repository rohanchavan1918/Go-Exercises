package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupModels() *gorm.DB {

	db, err := gorm.Open("mysql", "root:@/gotest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		panic("error db")
	}
	// defer db.Close()
	db.AutoMigrate(&User{})
	return db
}
