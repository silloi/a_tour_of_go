package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for i := 1; i <= 100; i++ {
		z1 := z
		z -= (z*z - x) / (2 * z)
		if z1 == z {
			return z
		}
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
