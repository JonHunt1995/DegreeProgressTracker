package main

import "fmt"

func main() {
	//ParseConfig can probably be moved to init() later
	c, e := parseConfig()

	if e != nil {
		panic(e)
	}
	fmt.Println(getDocsClient(*c))
}
