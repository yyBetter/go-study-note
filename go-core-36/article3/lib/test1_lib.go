package lib

import (
	"fmt"
	"go-study-note/go-core-36/article3/lib/internal"
	"os"
)

// 注意首字母大写，大写类似java中public关键字
func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func Hello1(name string) {
	internal.Hello(os.Stdout, name)
}
