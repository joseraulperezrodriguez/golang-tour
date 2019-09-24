package flowcontrol

import (
	"fmt"
)

//MainControlFlow is the main function for control flow tests
func MainControlFlow() {
	fmt.Println("Start MainControlFlow >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	ClassicFor()
	WhileFor()
	If()
	IfVar()
	SwitchBasic()
	Switch1()
	SwitchNoCondition()
	Defer()
	DeferStack()
	fmt.Println("End MainControlFlow <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println()
}
