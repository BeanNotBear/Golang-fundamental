package main

// file scope
import "fmt"

// package scope
const a = 1

func main() {
	fmt.Println(a)
	fmt.Println("123")

	if a == 1 {
		// block  scope
		b := 1
		fmt.Println(b)
	}
}
