package main

import "fmt"

func main() {
	var n int
	fmt.Print("please input your age : ")
	_ , err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)
	result := isPrimeNumber(n)
	if result {
		fmt.Println("number is prime")
	} else {
		fmt.Println("number is not prime")
	}

}

func isPrimeNumber(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
