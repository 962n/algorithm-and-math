package main

import (
	"fmt"
	"testing"
)

func TestGetMaxCommonDivisor(t *testing.T) {
	data := []struct {
		X      int
		Y      int
		Result int
	}{
		{X: 372, Y: 506, Result: 2},
		{X: 24, Y: 40, Result: 8},
	}

	for index, v := range data {
		actual := Gcd(v.X, v.Y)
		if actual != v.Result {
			t.Error(fmt.Printf("%d test error!!!", index))
		}
	}
}
