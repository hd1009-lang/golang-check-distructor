package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load("./envs/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
