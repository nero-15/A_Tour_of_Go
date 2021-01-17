// https://go-tour-jp.appspot.com/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	res := 0.0
	for {
		res = z - (z*z-x)/(2*z)
		if math.Abs(res-z) < 0.001 {
			break
		}
		z = res
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
