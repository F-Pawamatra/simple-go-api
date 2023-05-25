package database

import (
	"fmt"

	"github.com/f-pawamatra/rest-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root@tcp(localhost:3306)/go_rest?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database: %s", err.Error())
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	return db
}
