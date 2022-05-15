package main

import (
	"errors"
	"fmt"
)

func main() {
	var n int
	fmt.Print("please input number : ")
	_, scanErr := fmt.Scan(&n)
	if scanErr != nil {
		fmt.Println(scanErr.Error())
		return
	}
	var result []int
	tmp := n
	for {
		min, err := getMinDivisor(tmp)
		if err != nil {
			result = append(result, tmp)
			break
		}
		result = append(result, min)
		tmp = tmp / min
	}
	for _, v := range result {
		fmt.Println(v)
	}
}

func getMinDivisor(n int) (int, error) {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return i, nil
		}
	}
	return 0, errors.New("nothing")
}
