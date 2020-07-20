package main

import "fmt"

const gc = "gc"

const (
	n1 = iota // 0
	n2 // 1
	n3 = iota // 2
	n4 // 3
	n5 = 1 // 1
	n6 // 1
	n7 = iota // 6
)

func main() {
	var name string
	name = "gc"

	var age int
	fmt.Println("hello world", name, age)
	fmt.Println(gc)
	fmt.Println(n1, n2, n3, n4, n5, n6, n7)
}
