package main

import "go-study-note/go-core-36/article3/lib"

import "flag"

func main() {
	name := flag.String("name", "n", "enter your name.")
	flag.Parse()
	lib.Hello(*name)
	lib.Hello1(*name)
}
