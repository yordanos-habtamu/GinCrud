package config


import (
	"log"
	

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// You can load environment variables here, e.g.:
	// os.Getenv("PORT")
}
