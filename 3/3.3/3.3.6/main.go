package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(200000)
	if n < 2 {
		n = 2
	}
	fmt.Println(fmt.Sprintf("n = %d", n))
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		a := rand.Intn(99999)
		if a < 1 {
			a = 1
		}
		cards[i] = a
	}
	fmt.Println("dummy data created")

	count := map[int]int{}

	for _, v := range cards {
		_, ok := count[v]
		if !ok {
			count[v] = 0
		}
		count[v] += 1
	}
	fmt.Println("counting cards finish")

	result := 0
	for k, v := range count {
		if k > 50000 {
			continue
		}
		comp := 100000 - k
		cc, ok := count[comp]
		if !ok {
			continue
		}
		if v == 50000 {
			result = result + (v*(v-1)/2)
			continue
		}
		result = result + (v * cc)
	}
	fmt.Println(fmt.Sprintf("result : %d", result))

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
