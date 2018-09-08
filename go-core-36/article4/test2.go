package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v\n", *name)
}

func getTheFlag() *string {
	return flag.String("name", "w", "enter your name.")
}
