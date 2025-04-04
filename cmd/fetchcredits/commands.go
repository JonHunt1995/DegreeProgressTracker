package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JonHunt1995/DegreeProgressTracker.git/internal/config"
)

func commandFetchCompletedCUs(cfg config.Config) (int, error) {
	numCredits, err := config.GetGoogleSheet(cfg)
	if err != nil {
		return 0, err
	}
	return numCredits, nil
}

func commandHelp() error {
	fmt.Println("Hello, here are the available commands:")
	fmt.Println("\tcredits: Get the completed CUs at WGU")
	fmt.Println("\thelp: The manual for this CLI")
	fmt.Println("\texit: Get out of this program")
	return nil
}
func commandExit() error {
	fmt.Println("Thanks for using this application!")
	os.Exit(0)
	return nil
}

func executeCommand(cfg config.Config, commandName string) error {
	switch commandName {
	case "credits":
		numCUs, err := commandFetchCompletedCUs(cfg)
		if err != nil {
			return err
		}
		fmt.Println(numCUs)

	case "exit":
		if err := commandExit(); err != nil {
			return err
		}

	case "shell":
		runREPL(cfg)
	default:
		if err := commandHelp(); err != nil {
			return err
		}
	}
	return nil
}

func runREPL(cfg config.Config) error {
	r := os.Stdin
	s := bufio.NewScanner(r)
	// Runs a Loop For Commands Until Prompted To Exit
	for {
		fmt.Print("DegreeProgressTracker > ")
		// Scan for Input
		if ok := s.Scan(); !ok {
			fmt.Println("Unable to find command name")
			continue
		}
		// Sanitize and Validate Input
		inputs := strings.TrimSpace(s.Text())
		words := sanitizeInputs(inputs)
		if len(words) < 1 {
			fmt.Println("Unable to find command name")
			continue
		}
		commandName := words[0]

		// Execute Command
		if err := executeCommand(cfg, commandName); err != nil {
			return err
		}
	}
}
