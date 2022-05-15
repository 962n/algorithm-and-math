package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// random dummy data
	n := rand.Intn(200000)
	if n < 2 {
		n = 2
	}
	prices := []int{100, 200, 300, 400}
	goods := make([]int, n, n)
	for i := 0; i < n; i++ {
		index := rand.Intn(200000) % len(prices)
		goods[i] = prices[index]
	}

	// dummy data
	goods = []int{
		100, 100, 100, 100,
		200, 200,
		300, 300, 300,
		400, 400, 400, 400,
	}

	// action
	m := map[int]int{
		100: 0,
		200: 0,
		300: 0,
		400: 0,
	}
	for _, v := range goods {
		m[v] += 1
	}

	fmt.Println(m[100]*m[400] + m[200]*m[300])

}
