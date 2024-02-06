package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for nz := z - (z*z-x)/(2*z); nz-z > 1.0e-14 || z-nz > 1.0e-14; nz = z - (z*z-x)/(2*z) {
		z = nz
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
