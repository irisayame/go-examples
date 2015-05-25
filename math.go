package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))
	fmt.Printf("PI:%f\n", math.Pi)
	var z complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("%T %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("%T %v\n", f, f)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
