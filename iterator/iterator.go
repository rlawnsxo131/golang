package main

import "fmt"

func iterator() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}

func iteratorSecond() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func iteratorThird(v int) int {
	if custom := 10; v < custom {
		return v
	} else {
		return 10
	}
}

func main() {
	iterator()
	iteratorSecond()
	iteratorThird(3)
}
