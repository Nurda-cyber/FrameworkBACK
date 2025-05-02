package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JwtSecret string

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env файлы табылмады. JWT_SECRET бос болуы мүмкін.")
	}
	JwtSecret = os.Getenv("JWT_SECRET")

	dsn := "host=localhost user=postgres password=postgres dbname=Framework port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}
