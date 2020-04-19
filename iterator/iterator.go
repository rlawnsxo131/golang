package iterator

import "fmt"

func Iterator() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}

func IteratorSecond() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func IteratorThird(v int) int {
	if custom := 10; v < custom {
		return v
	} else {
		return 10
	}
}
