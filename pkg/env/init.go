package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init(path ...string) {
	if len(path) == 0 {
		path[0] = ".env"
	}

	err := godotenv.Load(path[0])
	if err != nil {
		panic("Error loading .env file")
	}

	if os.Getenv("DEBUG") == "true" {
		log.Print("Env initialized")

	}
}
