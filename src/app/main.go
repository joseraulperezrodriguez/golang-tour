package main

import (
	"golang-tour/src/arrays"
	"golang-tour/src/errors"
	"golang-tour/src/flowcontrol"
	"golang-tour/src/functional"
	"golang-tour/src/functions"
	"golang-tour/src/goroutines"
	"golang-tour/src/interfaces"
	"golang-tour/src/io"
	"golang-tour/src/maps"
	"golang-tour/src/methods"
	"golang-tour/src/pointers"
	"golang-tour/src/ranges"
	"golang-tour/src/structs"
	"golang-tour/src/types"
	"golang-tour/src/variables"
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
