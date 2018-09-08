package main

import (
	"fmt"
	"io"
	"os"
)

// go是由各种不同作用范围的代码快组成的，可以对一个参数设置不同类型的值，在不同代码块中；
// 如果是相同代码块，那么类型必须相同
func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	{
		err := 1
		fmt.Println(err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
}
