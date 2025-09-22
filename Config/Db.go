package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ExpencesManagment/Models"
)

func DatabaseConnection() *gorm.DB{
	dsn := "root:RevDau@tcp(127.0.0.1:3306)/expense?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err !=nil {
		fmt.Println("Database not connected")
	}

	db.AutoMigrate(&models.Expense{})
	return db
}