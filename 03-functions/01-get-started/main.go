package main

import "fmt"

// how to declare a function in golang?
//
//	func name(parameter datatype) {
//		return value
//	}

// example
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
