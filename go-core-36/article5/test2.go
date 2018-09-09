package main

import (
	"fmt"
	//. "go-study-note/go-core-36/article5/lib" 如果引入包名称重复，会报错，因为在一个作用域下
	"go-study-note/go-core-36/article5/lib"
)

var T = "main"

func main() {
	fmt.Println(lib.T)
}
