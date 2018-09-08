// 模块级私有，代码包中声明的公开程序实体仅能被该代码包的直接父包及其子包中的代码引用
package internal

import (
	"fmt"
	"io"
)

func Hello(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s!\n", name)
}
