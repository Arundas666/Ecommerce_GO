package database

import (
	"firstpro/config"
	"firstpro/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	pasqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dberr := gorm.Open(postgres.Open(pasqlInfo), &gorm.Config{})
	if dberr != nil {
		return nil, fmt.Errorf("failed 😁 to connect to database:%w", dberr)
	}

	DB = db
	DB.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Admin{})
	
	return DB, dberr

}
