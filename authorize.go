// Authorize will be used to handle the authorization of sessions between the Google Docs API and the Discord API
// At this time, only the Google Docs portion is being implemented. My intention is that the Google Docs portion of
// this file will provide a fully useable set of data from the spreadsheet.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
	"os"
)

func getDocsData() {
	c, e := parseConfig()
	fmt.Print("Using credentials:\n", c.String())
	if e != nil {
		panic(e)
	}
	client := getDocsClient(c)
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	//srv, err := sheets.NewService(context.Background(), option.WithAPIKey(c.apiKey))
	if err != nil {
		log.Fatalf("Unable to start a Sheets service: %v", err)
	}

	// Service -> SpreadsheetsService -> SpreadsheetsValuesService -> Get -> SpreadsheetsValuesGetCall -> ValueRange
	resp, err := srv.Spreadsheets.Values.Get(c.SpreadsheetID, c.readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// resp.Values is a [][]interface{}
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("%v x %v", len(resp.Values), len(resp.Values[0]))
	}
}

// getDocsClient loads uses a configuration struct to find am auth token from a file. It then returns a pointer to an
// oauth2 token
func getDocsClient(c *Config) *http.Client {
	path := fmt.Sprintf("./config/%s", c.docsTokenPath)
	fmt.Println(path)
	tok := getDocsToken(path)
	conf := getDocsConfig(path)
	fmt.Println("Authorized with Google Docs")

	return conf.Client(context.Background(), tok)
}

// getDocsConfig takes a file path and returns an oauth2 config object. It handles its own errors
func getDocsConfig(path string) *oauth2.Config {
	b, err := os.ReadFile(path)
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file and create config: %v", err)
	}
	return config
}

// getDocsToken takes a file path and returns an oauth2 token. It handles its own errors
func getDocsToken(path string) *oauth2.Token {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to parse client secret file and create token: %v", err)
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		fmt.Println("Error decoding json: ", err)
	}
	fmt.Println(tok.AccessToken)
	return tok
}
