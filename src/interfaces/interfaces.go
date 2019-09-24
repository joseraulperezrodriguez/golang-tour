package interfaces

import (
	"fmt"
	"math"
)

//Abser interface definition
type Abser interface {
	Abs() float64
}

//Any type match 'Any' interface
type Any interface{}

//MyFloat custom type
type MyFloat float64

//Vertex type
type Vertex struct {
	X, Y float64
}

//Abs function
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Abs function
func (mf MyFloat) Abs() float64 {
	if mf < 0 {
		return float64(mf) * -1
	}
	return float64(mf)
}

//Interface1 test
func Interface1() {
	var absI Abser
	var mf = MyFloat(5.1)
	var v = Vertex{0, 1}

	absI = mf
	absI = &v

	fmt.Println(absI)

	describe(absI)
	describe(mf)
	describe(&v)
}

//Interface2 test
func Interface2() {
	var absI Abser = MyFloat(5.0)
	ta, ok := absI.(MyFloat)
	println(ta, ok)
}

//Interface3 test
func Interface3() {
	var absI Abser = MyFloat(5.0)

	switch typ := absI.(type) {
	case MyFloat:
		fmt.Println("MyFloat", typ)
	case *Vertex:
		fmt.Println("Vertex", typ)
	default:
		fmt.Println("Other", typ)
	}
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}
