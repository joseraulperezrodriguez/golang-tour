package functions

import "fmt"

//Add learing function
func Add(x int, y int) int {
	return add(x, y)
}

func add(x, y int) int {
	return x + y
}

//Swap function effectively swap two element in array
func Swap(a, b string) (string, string) {
	return b, a
}

//NamedReturn is a function to learning using named return parameter
func NamedReturn(a bool) (b bool) {
	b = !a
	return
}

//Variadic test
func Variadic(array ...int) {
	fmt.Println(array)
}
