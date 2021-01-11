package structs

import "fmt"

//Vertex struct sample
type Vertex struct {
	X, Y int
}

//CreateVertex function
func CreateVertex() {
	v := Vertex{3, 4}
	fmt.Println(Vertex{1, 2})
	u := &v
	fmt.Println(v.X, v.Y)
	fmt.Println((*u).X, (*u).Y)
	var vertex *Vertex = &(*u)
	fmt.Println((*vertex).X, (*vertex).Y)
}

//StructLiterals show that zero values are used for not defined fields
func StructLiterals() {
	v := Vertex{X: 0, Y: 0}
	fmt.Println(v)
}
