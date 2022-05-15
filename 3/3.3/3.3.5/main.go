package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	// random dummy data
	n := rand.Intn(500000)
	if n < 2 {
		n = 2
	}
	colors := []int{1, 2, 3} // red , blue , yellow
	cards := make([]int, n, n)
	for i := 0; i < n; i++ {
		index := rand.Intn(200000) % len(colors)
		cards[i] = colors[index]
	}

	// dummy data
	cards = []int{
		1, 1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3,
	}

	// action
	m1 := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}
	for _, v := range cards {
		m1[v] += 1
	}
	m2 := map[string]int{
		"red":    1,
		"blue":   2,
		"yellow": 3,
	}
	for k, v := range m2 {
		if m1[v] < 2 {
			fmt.Println(k + ": noting")
			continue
		}
		res, err := C(m1[v], 2)
		if err != nil {
			fmt.Println(k + ": error!")
			continue
		}
		fmt.Println(fmt.Sprintf("%s:%d", k, res))
	}

}

func C(n, r int) (int, error) {
	c, err := P(n, r)
	if err != nil {
		return 0, err
	}

	m, err := P(r, r)
	if err != nil {
		return 0, err
	}

	return c / m, nil
}

func P(n, r int) (int, error) {
	if n-r < 0 {
		return 0, errors.New("parameter invalid")
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= n - i
	}
	return result, nil
}
