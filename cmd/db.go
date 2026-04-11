package main

import (
	"fmt"
	"template-go/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dbHost := "postgres"
	dbPort := "5431"
	dbName := "database"
	dbUser := "shuneo"
	dbPass := "123123123"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", dbHost, dbUser, dbPass, dbName, dbPort)
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	user := &model.User{}
	user.BeforeCreate()
	_ = db.AutoMigrate(user)

}
