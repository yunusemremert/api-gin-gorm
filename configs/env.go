package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvPostgreSQLDBURL() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
	}

	postgreSQLUrl := os.Getenv("POSTGRESQL_URL")
	if postgreSQLUrl == "" {
		log.Fatal("You must set your 'POSTGRESQL_URL' environmental variable.")
	}

	return postgreSQLUrl
}
