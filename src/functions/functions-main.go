package functions

import (
	"fmt"
)

//MainFunctions is the main function for function package
func MainFunctions() {
	fmt.Println("Start MainFunctions >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(Add(2, 5))
	fmt.Println(Swap("world", "hello"))
	fmt.Println(NamedReturn(false))
	fmt.Println("End MainFunctions <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println()
}
