package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env vars", err.Error())
		os.Exit(1)
	}
}
