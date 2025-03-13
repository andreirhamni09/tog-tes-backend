package database

import (
	"backend/models/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// docker run --name my_postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=tog -p 5432:5432 -d postgres
	godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal Terkoneksi Ke Database : " + err.Error())
	}
	DB = db
	log.Println("Berhasil Terkoneksi Kedatabase")

	// Auto Migrate
	db.AutoMigrate(&entity.Customer{})
	db.AutoMigrate(&entity.BankAccount{})
	db.AutoMigrate(&entity.PocketInformation{})
	db.AutoMigrate(&entity.Deposito{})
	db.AutoMigrate(&entity.DepositoHistory{})
}
