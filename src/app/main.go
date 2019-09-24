package main

import (
	"functions"
	"tour/src/arrays"
	"tour/src/errors"
	"tour/src/flowcontrol"
	"tour/src/functional"
	"tour/src/goroutines"
	"tour/src/interfaces"
	"tour/src/io"
	"tour/src/maps"
	"tour/src/methods"
	"tour/src/pointers"
	"tour/src/ranges"
	"tour/src/structs"
	"tour/src/types"
	"variables"
)

func main() {
	functions.MainFunctions()
	variables.MainVariables()
	types.MainTypes()
	flowcontrol.MainControlFlow()
	pointers.MainPointers()
	structs.MainStructs()
	arrays.MainArrays()
	ranges.MainRanges()
	maps.MainMaps()
	functional.MainFunctional()
	methods.MainMethods()
	interfaces.MainInterfaces()
	errors.MainErrors()
	io.MainIO()
	goroutines.MainGoRoutines()
}
