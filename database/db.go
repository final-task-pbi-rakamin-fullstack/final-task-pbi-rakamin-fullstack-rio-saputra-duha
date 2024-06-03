package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func DBConnect() {
	var err error
	dsn := os.Getenv("DB_URL")
	//fmt.Println(dsn)
	//fmt.Println(os.Getenv("SECRET_KEY"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//Created By Rafly Andrian
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}
