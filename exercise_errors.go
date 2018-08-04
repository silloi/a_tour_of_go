package main
// wakarahenn
import (
	"fmt"
)

type ErrNegativeSqrt (float64, erorr) {
	
}

func Sqrt(x float64) float64 {
	if x < 0 {
		 return ErrNegativeSqrt()
	}
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

func (e ErrNegativeSqrt) Error() string {
	fmt.Println("cannot Sqrt negative number:",
	float64(e))
}

//func Sqrt(x float64) (float64, error) {
//	return 0, nil
//}

func main() {
	if err :=run(); err != nil {
		Error()
	}
}
