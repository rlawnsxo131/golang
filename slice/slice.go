package slice

import "fmt"

func Slice() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	// [3, 5, 7]
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	// [2, 3, 5]
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	// [11, 13]
	fmt.Println("p[4:] ==", p[4:])
}

func SliceSecond() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5) // capacity
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func SliceThird() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil")
	}
}

// range
func SliceFourth() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func SliceFifth() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	// ignore index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
