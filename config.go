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
// the config path could probably be moved to a global variable, but I don't think it is necessary
func parseConfig() (*Config, error) {
	f, err := os.OpenFile("./config/config.txt", os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	//Convert the file to a scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	//Initialize a new Config object. Later this can be filled with default values.
	config := Config{}

	//Here we can iterate through the config file, line by line.
	//Using a switch, initialize the config object with variables
	for scanner.Scan() {
		//Split each line by a "=" because I learned Spring once
		tokens := strings.Split(scanner.Text(), "=")
		switch tokens[0] {
		case "docs-token":
			config.docsTokenPath = tokens[1]
		case "docs-id":
			config.SpreadsheetID = tokens[1]
		}
	}
	//fmt.Println(config)
	return &config, nil
}

func (c *Config) String() string {
	return fmt.Sprintf("docs-token: %s\ndocs-id: %s\n", c.docsTokenPath, c.SpreadsheetID)
}
