package main

import (
	"golang-tour/arrays"
	"golang-tour/errors"
	"golang-tour/flowcontrol"
	"golang-tour/functional"
	"golang-tour/functions"
	"golang-tour/goroutines"
	"golang-tour/interfaces"
	"golang-tour/io"
	"golang-tour/maps"
	"golang-tour/methods"
	"golang-tour/pointers"
	"golang-tour/ranges"
	"golang-tour/structs"
	"golang-tour/types"
	"golang-tour/variables"
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
