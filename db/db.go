package db

import (
	"fmt"

	"realtime-chat/config"
	"realtime-chat/modules/user/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrasi User Model
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
