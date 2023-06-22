package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBinstance() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	log.Println("DB_HOST:", dbHost)
	dsn := "host=" + dbHost + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " port=5432 sslmode=disable "
	// dsn := "host=localhost  user=postgres  dbname=database  password=123 port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Database connection established")
	DB = db
	return db, nil
}
