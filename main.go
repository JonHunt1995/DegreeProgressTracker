package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := os.Stdin
	s := bufio.NewScanner(r)
	// Runs a Loop For Commands Until Prompted To Exit
	for {
		fmt.Print("DegreeProgressTracker > ")
		if ok := s.Scan(); !ok {
			break
		}
	}
}
