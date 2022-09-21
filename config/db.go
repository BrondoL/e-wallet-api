package config

import (
	"fmt"
	"log"
	"os"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.PasswordReset{}, &model.Wallet{}, &model.SourceOfFund{}, &model.Transaction{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetConn() *gorm.DB {
	return db
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
