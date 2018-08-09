package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z1 := z
	z2 := z1
	for i := 0; ; i++ {
		z2 = z1
		z1 = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		if z == z1 || z == z2 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
