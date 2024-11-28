package main

import "fmt"

func main() {
	i, j := 1, 2

	p := &i
	fmt.Println(*p)

	p = &j
	*p = *p * 5
	fmt.Println(*p)
}
