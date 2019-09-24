package functional

import (
	"fmt"
	"math"
)

//Functional1 tests

func compute(function func(float64, float64) float64) float64 {
	return function(3, 4)
}

//Functional1 tests
func Functional1() {
	var hypot = func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(2, 5))
	fmt.Println(compute(hypot))
}

func adder() func(int) int {
	var num = 1

	return func(a int) int {
		return num + a
	}

}

//Closure show closures in Golang
func Closure() {
	fmt.Println(adder()(2))
}
