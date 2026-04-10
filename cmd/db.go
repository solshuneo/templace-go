package main

import (
	"fmt"
	"lotesaleagent/internal/repository/gsql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "postgres"
	dbUser := "shuneo"
	dbPass := "123123123"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", dbHost, dbUser, dbPass, dbName, dbPort)
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = db.AutoMigrate(&gsql.User{})

}
