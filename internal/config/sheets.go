package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetGoogleSheet(cfg Config) (int, error) {
	// Auth Logic to Connect to Spreadsheet
	ctx := context.Background()
	jsonKey, err := os.ReadFile(cfg.GoogleSheetsJSON)
	if err != nil {
		return 0, fmt.Errorf("Unable to read credentials file: %v", err)
	}

	srv, err := sheets.NewService(ctx, option.WithCredentialsJSON(jsonKey))
	if err != nil {
		return 0, fmt.Errorf("Unable to retrieve Sheets client: %v", err)
	}
	// Fetch Data From Specified Spreadsheet
	resp, err := srv.Spreadsheets.Values.Get(cfg.SpreadsheetID, cfg.TotalCreditsRange).Do()
	if err != nil {
		return 0, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}
	if len(resp.Values) == 0 {
		return 0, fmt.Errorf("No data found")
	}
	// Return Data
	numCredits, err := strconv.Atoi(resp.Values[0][0].(string))
	if err != nil {
		return 0, fmt.Errorf("Error in converting resp value to an int")
	}
	return numCredits, nil
}