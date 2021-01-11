package types

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	//ToBe variable
	ToBe bool = false
	//MaxInt variable
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//PrintTypes prints variables
func PrintTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

//TypeConvertions show how to convert between types
func TypeConvertions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

//TypeInference show how type inference works
func TypeInference() {
	v := 42 + 5i // change me!
	fmt.Printf("v is of type %T\n", v)
}
