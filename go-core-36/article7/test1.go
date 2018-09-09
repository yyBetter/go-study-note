package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s2 := make([]int, 5, 8)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	s5 := s4[:cap(s4)]
	fmt.Println(len(s5))
	fmt.Println(cap(s5))
}
