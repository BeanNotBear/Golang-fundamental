package main

import "fmt"

// x int, y int can be written as x, y int
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
