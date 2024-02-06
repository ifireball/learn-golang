package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for nz := z - (z*z-x)/(2*z); nz-z > 1.0e-14 || z-nz > 1.0e-14; nz = z - (z*z-x)/(2*z) {
		z = nz
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
