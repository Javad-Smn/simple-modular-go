package libs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DotEnvVariable -> get .env
func DotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(err.Error())
	}

	return os.Getenv(key)
}
