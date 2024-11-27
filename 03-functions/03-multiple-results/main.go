package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	x, y := swap(1, 2)
	fmt.Println(x, y)
}
