package main

import "fmt"

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func addrMain() {
	pos, neg := addr(), addr()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func fibonacci() func() int {
	val := 0
	val1 := 0
	return func() int {
		val2 := val1 + val
		val = val1
		val1 = val2
		if val1 == 0 {
			val++
		}
		return val2
	}
}

func fibonacciMain() {
	fibo := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibo())
	}
}

func main() {
	addrMain()
	fibonacciMain()
}
