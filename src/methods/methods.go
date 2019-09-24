package methods

import (
	"fmt"
	"math"
)

//Vertex struct
type Vertex struct {
	X, Y float64
}

/*//Abs testing
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y + v.Y)
}*/

//Abs testing
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y + v.Y)
}

//Increase testing
func (v *Vertex) Increase() {
	v.X++
	v.Y++
}

//MyInt16 type
type MyInt16 int16

//Abs on int16 numbers
func (i MyInt16) Abs() MyInt16 {
	if i > 0 {
		return i
	}
	return i * -1
}

//Method1 test
func Method1() {
	v := Vertex{1, 2}
	u := &v
	fmt.Println(v.Abs())
	u.Increase()
	fmt.Println(u)
	v.Increase()
	fmt.Println(v)
}

//Method2 test
func Method2() {
	var i MyInt16 = 18
	var j int16 = 19
	fmt.Println(i.Abs(), (MyInt16(j)).Abs())
}
