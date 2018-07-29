package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	sum := 0
	sum_k_1 := 0
	sum_k_2 := 0
	return func(x int) int {
		sum_k_2 = sum_k_1
		sum_k_1 = sum
		sum = sum_k_1 + sum_k_2
		if x == 1 {
			sum = 1
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}