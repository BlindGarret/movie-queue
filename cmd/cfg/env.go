package cfg

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFiles() {
	// Load dev env file
	if _, err := os.Stat(".env.dev"); !errors.Is(err, os.ErrNotExist) {
		err := godotenv.Load(".env.dev")
		if err != nil {
			panic("Error loading .env file")
		}
	}

	// Load base env file
	if _, err := os.Stat(".env"); !errors.Is(err, os.ErrNotExist) {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}

}
