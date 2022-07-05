package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("local:@tcp(127.0.0.1:3306)/perpus?charset=utf8mb4&parseTime=True"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
