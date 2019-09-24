package variables

import (
	"fmt"
)

var i, j int

var ii, bb = 1, "bb"

//PrintVars print default variables for int
func PrintVars() {
	fmt.Println(i, j)
	fmt.Println(ii, bb)
	k := true
	fmt.Println(k)
}

//VariableFormat function
func VariableFormat() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

//Constants uses constant values
func Constants() {
	const World = "世界"
	fmt.Println("Hello", World)
}

const (
	//Big Create a huge number by shifting a 1 bit left 100 places.
	//Big In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	//Small Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int64) int64 { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

//AdjustConstant test
func AdjustConstant() {
	needFloat(Big)
}
