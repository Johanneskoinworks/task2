package database

import (
	"fmt"
	"log"
	"task2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Start(host string, user string, password string, dbname string, dbPort string) {
	// config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user,
	// 	password, dbname)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}
	db.Debug().AutoMigrate(models.Product{}, models.Order{})
}
func GetDB() *gorm.DB {
	return db
}
