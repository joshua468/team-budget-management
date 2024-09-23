package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joshua468/team-budget-management/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("PORT"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database:%v", err)
	}
	fmt.Println("Database connected successfully!")
	DB.AutoMigrate(&models.Team{}, &models.Expense{})
}
