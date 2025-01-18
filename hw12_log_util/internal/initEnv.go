package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() (string, string, string) {
	if err := LoadEnv(); err != nil {
		log.Fatalf("load env variables: %s", err.Error())
	}
	inputFile := os.Getenv("LOG_ANALYZER_FILE")
	verbose := os.Getenv("LOG_ANALYZER_LEVEL")
	outputFile := os.Getenv("LOG_ANALYZER_OUTPUT")
	return inputFile, verbose, outputFile
}

func LoadEnv() error {
	configPath := ".env"
	if err := godotenv.Load(configPath); err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	return nil
}
