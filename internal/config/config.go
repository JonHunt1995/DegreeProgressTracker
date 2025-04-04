package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleSheetsJSON  string
	SpreadsheetID     string
	TotalCreditsRange string
	DiscordID         string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		GoogleSheetsJSON:  os.Getenv("GOOGLE_SHEETS_JSON"),
		SpreadsheetID:     os.Getenv("SPREADSHEET_ID"),
		TotalCreditsRange: os.Getenv("TOTAL_CREDITS_RANGE"),
		DiscordID:         os.Getenv("ALLOWED_USER_ID"),
	}, nil
}
