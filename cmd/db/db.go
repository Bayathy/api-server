package db

import (
	"fmt"
	"github.com/bayathy/api-server/cmd/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDatabase() (*gorm.DB, error) {
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("postgresql://bayathy:%s@dune-gorilla-4292.6xw.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full", password)
	log.Println(password)
	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = db.AutoMigrate(&entity.Article{}, &entity.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
