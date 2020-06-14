package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func makeMap() {
	var m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func makeMapSecond() {
	var m = map[string]Vertex{
		// Vertex 생략 가능
		// 만약 가장 상위의 타입이 타입명이라면 리터럴에서 타입명을 생략해도 됩니다.
		// "Bell Labs": Vertex{
		// 	40.68433, -74.39967,
		// },
		// "Google": Vertex{
		// 	37.42202, -122.08408,
		// },
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func makeMapThird() {
	m := make(map[string]int)

	// 맵의 요소 삽입, 수정, 가져오기
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 맵 요소 제거
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 키 존재여부 확인
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

func main() {
	makeMap()
	makeMapSecond()
	makeMapThird()
}
