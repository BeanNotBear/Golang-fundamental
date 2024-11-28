package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1 = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.Y)

	// pointer to struct
	p := &v
	fmt.Println(*p)

	// when assign to this it changes to vertex v
	p.X = 3
	fmt.Println(p.X, p.Y)
	fmt.Println(v)

	fmt.Println(v1, p1, v2, v3)
}
