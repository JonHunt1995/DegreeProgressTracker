package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleSheetsJSON  string
	SpreadsheetID     string
	TotalCreditsRange string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		GoogleSheetsJSON: os.Getenv("GOOGLE_SHEETS_JSON"),
		SpreadsheetID: os.Getenv("SPREADSHEET_ID"),
		TotalCreditsRange: os.Getenv("TOTAL_CREDITS_RANGE"),
	}
}