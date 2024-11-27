package main

import "fmt"

func main() {
	defer fmt.Println("world") // run after the last action

	fmt.Println("hello")
}
