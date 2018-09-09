package main

import "fmt"

func main() {
	// 示例1。
	// 只能发不能收的通道。
	var uselessChan = make(chan<- int, 1)
	// 只能收不能发的通道。
	var anotherUselessChan = make(<-chan int, 1)
	// 这里打印的是可以分别代表两个通道的指针的16进制表示。
	fmt.Printf("The useless channels: %v, %v\n",
		uselessChan, anotherUselessChan)
}
