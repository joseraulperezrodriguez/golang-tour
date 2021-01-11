package maps

import "fmt"

//Vertex struct
type Vertex struct {
	X, Y int32
}

//Map1 test functions
func Map1() {
	var map1 map[string]Vertex = make(map[string]Vertex)

	map1["Center"] = Vertex{0, 0}

	var v, err = map1["Corner"]

	fmt.Println(map1)
	if err {
		fmt.Println("Not found")
	} else {
		fmt.Println(v)
	}
}

//Map2 test literal
func Map2() {
	var map1 = map[string]Vertex{
		"One":   {51, 3},
		"Two":   Vertex{1, 31},
		"Three": {1, 31},
	}
	fmt.Println(map1)
}

//Map3 for more map examples
func Map3() {
	var map1 = make(map[string]int)
	map1["k"] = 1
	delete(map1, "k")
}
