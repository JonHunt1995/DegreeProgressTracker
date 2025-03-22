package main

import "fmt"

func main() {
	//r := os.Stdin
	//s := bufio.NewScanner(r)
	c, e := parseConfig()

	if e != nil {
		panic(e)
	}
	fmt.Println(c)
}
