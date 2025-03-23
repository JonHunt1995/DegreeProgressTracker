package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JonHunt1995/DegreeProgressTracker.git/internal/config"
)

func commandFetchCompletedCUs(cfg config.Config) (int, error) {
	fmt.Printf("Fetching from %v\n", cfg.SpreadsheetID)
	return 60, nil
}

func commandHelp(cfg config.Config) error {
	fmt.Println("Hello, here are the available commands:")
	fmt.Println("\tfetchCUs: Get the completed CUs at WGU")
	fmt.Println("\thelp: The manual for this CLI")
	fmt.Println("\tfexit: Get out of this program")
}
func commandExit(cfg config.Config) error {
	fmt.Println("Thanks for using this application!")
	os.Exit(0)
	return nil
}

func executeCommand(cfg config.Config, commandName string) {
	switch commandName {
	case "fetchCUs":
		commandFetchCompletedCUs(cfg)
	case "exit":
		commandExit(cfg)
	default:
		commandHelp(cfg)
	}
}
func runREPL() {
	cfg := config.LoadConfig()
	// Runs a Loop For Commands Until Prompted To Exit
	for {
		s := bufio.NewScanner(os.Stdin)
		fmt.Print("DegreeProgressTracker > ")
		if ok := s.Scan(); !ok {
			break
		}
		inputs := strings.TrimSpace(s.Text())
		words := sanitizeInputs(inputs)
		commandName := words[0]
		// Execute Command
		executeCommand(cfg, commandName)
	}
}
