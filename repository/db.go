package repository

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/Muhammad5943/GO-mini-ecommerce/model"
)

func DB() *gorm.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", username+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		log.Fatal("Error connecting to database")
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	return db
}
