package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JonHunt1995/DegreeProgressTracker.git/internal/config"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to log config due to %v", err)
	}

	if err = runREPL(cfg); err != nil {
		fmt.Println("No command found. Good luck")
		if err := commandHelp(); err != nil {
			fmt.Println("Error occured when prompted with help")
		}
	}
	os.Exit(0)
}
