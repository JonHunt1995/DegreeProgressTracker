package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	docsTokenPath string
	SpreadsheetID string
	test1         string
	test2         string
	test3         string
}

// parseConfig reads a hardcoded config file path so that various program parameters can be customized
func parseConfig() (*Config, error) {
	f, err := os.OpenFile("./config/config.txt", os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	//Convert the file to a scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	//here we can iterate through the config file, line by line.
	//Using a switch, initialize the config object with variables
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		tokens := strings.Split(scanner.Text(), "=")
		fmt.Println(tokens[0], tokens[1])
	}

	return &Config{}, nil
}
