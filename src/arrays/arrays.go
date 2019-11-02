package arrays

import "fmt"

//ArrayBasic array basic
func ArrayBasic() {
	var a [2]string
	a[0] = "A"
	a[1] = "rray"
	fmt.Println(a)

	b := [3]int{0, 1, 2}
	fmt.Printf("%v %T\n", b, b)

	var c []float64 = []float64{0.1, 0.1, 2.1}
	fmt.Printf("%v %T\n", c, c)

	d := append(c, 3.14)
	fmt.Println(d, len(d))
	fmt.Printf("%T\n", d)

}

//SlicesBasic examples
func SlicesBasic() {
	var arrayDefault []int = []int{2, 3, 5, 7, 11}
	var slice []int = arrayDefault[0:len(arrayDefault)]
	fmt.Println(slice)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

}

//Slice1 slice1 examples
func Slice1() {
	q := []int{2, 3, 5, 7}
	fmt.Println(q)

	s := []struct {
		x, y int
	}{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(s)
}

//SliceDefault shows default values for slices
func SliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s1 := append(s, 1)
	fmt.Println(s)
	fmt.Println(s1)
}

//SliceLenCap tests
func SliceLenCap() {
	var array [3]int
	var slice []int = array[0 : len(array)-1]
	fmt.Println(array)
	fmt.Println(slice)
	var sliceApp = append(slice, 1)
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(sliceApp)
	fmt.Println(cap(slice), len(slice))
}

//NilSlice example
func NilSlice() {
	var s []int
	if s == nil {
		fmt.Println(s)
	}
}

//MakingSlices show how to build slices using make function
func MakingSlices() {
	var slice1 = make([]int, 5)
	fmt.Println(len(slice1), cap(slice1))

	var slice2 = make([]int, 1, 5)
	fmt.Println(len(slice2), cap(slice2))
}

//SliceOfSlice example
func SliceOfSlice() {
	board := [][]string{
		[]string{"a", "b"},
		[]string{"c", "d"},
	}

	fmt.Println(board)
}
