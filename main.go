package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// remember export name must be uppercase the first character
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(swap("a", "b"))
	fmt.Println(split(17))

	// declare variable
	var i, j int = 2, 3
	fmt.Println(i, j)

	k := 1 // like var a  = 5 in c#
	fmt.Println(k)

	// remember if -- uint -- it mean unsigned int (non-negative integer)
	// byte // alias for uint8
	//rune // alias for int32

	// run in last program
	defer fmt.Println("Last func")

	fmt.Println("Func")
}
