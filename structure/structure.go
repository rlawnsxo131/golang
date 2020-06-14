package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func structure() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}

func structureSecond() {
	var (
		p = Vertex{1, 2}  // has type Vertex
		q = &Vertex{1, 2} // has type *Vertex
		r = Vertex{X: 1}  // Y:0 is implicit
		s = Vertex{}      // X:0 and Y:0
	)
	fmt.Println(p, q, r, s)
}

func structureThird() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}

func main() {
	structure()
	structureSecond()
	structureThird()
}
