package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	if os.Getenv("RAILWAY_ENV") == "production" {
		fmt.Println("Production environment")
	} else {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dsn := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}
	fmt.Println("Connected to database")
}
