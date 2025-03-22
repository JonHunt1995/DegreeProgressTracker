package main

import (
	"bufio"
	"fmt"
	"os"
)

func commandFetchCompletedCUs(cfg config.Config) (int, error) {
	fmt.Printf("Fetching from %v\n", cfg.SpreadsheetID)
	return 60, nil
}

func runREPL() {
	// Runs a Loop For Commands Until Prompted To Exit
	for {
		s := bufio.NewScanner(os.Stdin)
		fmt.Print("DegreeProgressTracker > ")
		if ok := s.Scan(); !ok {
			break
		}
		inputs := strings.TrimSpace(s.Text())
		words := sanitizeInputs(inputs)

	}
}
