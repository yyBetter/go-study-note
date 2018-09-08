package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

// init 方法会在main方法之前执行，类似java中静态代码快
func init() {
	flag.StringVar(&name, "name", "c", "enter your name.")
}

func main() {
	// 可以自定义参数使用说明
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	fmt.Printf("enter name: %v \n", name)
}

// flag 可以接收string int float等参数；此外还可以使用val方法来自定义类型
