package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := P(10,3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	res, err = C(2, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	res, err = C(8, 5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
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
	if n - r < 0 {
		return 0, errors.New("parameter invalid")
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= n - i
	}
	return result, nil
}
