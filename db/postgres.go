package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pluserver/domain"
)

func InitDB() (*gorm.DB, error) {
	//dsn := "host=db user=user dbname=mydb port=5432"
	dsn := "host=localhost user=user dbname=mydb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&domain.User{}, &domain.Transaction{}, &domain.Account{})
	if err != nil {
		log.Panicf("failed to migrate schema %v", err)
	}

	return db, err
}
