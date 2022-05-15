package main

import "fmt"

func main() {

	var list = []int{24, 40, 60, 80, 90, 120}
	l := len(list)
	if l < 2 {
		fmt.Println("list size must be more than 2")
		return
	}
	res := list[0]
	for _, number := range list {
		res = Lcm(res, number)
	}
	fmt.Printf("leastest common multiple is %d", res)
}

func Lcm(n, m int) int {
	gcd := Gcd(n, m)
	return gcd * (n / gcd) * (m / gcd)
}

func Gcd(n, m int) int {
	if n <= 0 {
		return m
	}
	if m <= 0 {
		return n
	}
	for n > 0 && m > 0 {
		if n < m {
			m = m % n
		} else {
			n = n % m
		}
	}
	return max(n, m)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
