package database

import (
	"fmt"

	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {

	viper.SetConfigFile("../../.env")
	viper.ReadInConfig()

	log.Println(viper.Get("DB_USERNAME"))

	dbUsername := viper.Get("DB_USERNAME")
	dbPassword := viper.Get("DB_PASSWORD")
	dbHost := viper.Get("DB_HOST")
	dbDatabaseName := viper.Get("DB_DATABASE_NAME")
	dbPort := viper.Get("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUsername, dbDatabaseName, dbPassword)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return db, err
	}

	postgresDb, err := db.DB()
	if err != nil {
		return db, err
	}
	if err := postgresDb.Ping(); err != nil {
		return db, err
	}

	return db, nil

}
