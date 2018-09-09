package main

func main() {
	// 示例1。
	//var badMap1 = map[[]int]int{} // 这里会引发编译错误。
	//_ = badMap1

	// 示例2。
	//var badMap2 = map[interface{}]int{
	//	"1":      1,
	//	[]int{2}: 2, // 这里会引发panic。
	//	3:        3,
	//}
	//_ = badMap2
}
