package main

import "fmt"

func main() {
	var srcInt = int16(-255)
	fmt.Println(srcInt)
	fmt.Println(uint16(srcInt))
	fmt.Println(int8(srcInt))
}
