package model

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Redirect string    `json:"redirect" gorm:"not null"`
	Goly     string    `json:"goly" gorm:"unique;not null"`
	Clicked  uint64    `json:"clicked"`
	Random   bool      `json:"random"`
}

func Setup() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading dotenv file: ", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not found in the environment")
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}
