package db

import (
	"fmt"

	"Ecom/pkg/config"
	"Ecom/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(domain.Admin{})
	db.AutoMigrate(domain.User{})
	db.AutoMigrate(domain.Product{})
	db.AutoMigrate(domain.Inventory{})
	db.AutoMigrate(domain.Order{})
	return db, dbErr
}
