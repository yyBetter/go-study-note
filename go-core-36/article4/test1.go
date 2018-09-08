package main

import (
	"flag"
	"fmt"
)

func main() {
	// way 1
	//var name string
	//flag.StringVar(&name, "name", "w", "enter your name.")

	// way 2 flag.String 获取的是一个指向string的指针，*号获取所指的string
	name := *flag.String("name", "w", "enter your name.")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}
