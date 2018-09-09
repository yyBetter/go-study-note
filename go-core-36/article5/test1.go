package main

import "fmt"

var block = "package"

// block 在不同的作用域可以被重新赋值，类型可以不同
func main() {
	block := "function"
	{
		block := "inner"
		fmt.Println(block)
	}
	fmt.Println(block)
}
