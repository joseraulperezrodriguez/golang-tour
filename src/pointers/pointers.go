package pointers

import (
	"fmt"
)

//BasicPointer example
func BasicPointer() {
	var i *int
	fmt.Printf("i= %v\n", i)
	j := 5
	i = &j
	fmt.Printf("i=%v value=%v\n", i, *i)
	*i = 6
	fmt.Printf("i=%v j=%v\n", *i, j)
	fmt.Printf("i=%v\n", i)
}
