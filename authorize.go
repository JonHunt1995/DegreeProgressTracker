// Authorize will be used to handle the authorization of sessions between the Google Docs API and the Discord API
// At this time, only the Google Docs portion is being implemented
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
)

// getDocsClient loads uses a configuration struct to find am auth token from a file. It then returns a pointer to an
// oauth2 token
func getDocsClient(c Config) *http.Client {
	path := fmt.Sprintf("./config/%s", c.docsTokenPath)
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
	err = json.NewDecoder(f).Decode((tok))
	return tok
}
