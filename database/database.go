package database

import (
	"fmt"
	"log"

	"github.com/qahta0/tasitracker/config"
	"github.com/qahta0/tasitracker/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	host := config.Config("DB_HOST")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	dbname := config.Config("DB_NAME")
	port := config.Config("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbConnection.AutoMigrate(&model.Company{}, &model.TodayPoints{})
	return dbConnection
}
