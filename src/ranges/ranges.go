package ranges

import "fmt"

//RangeBasic starting ranges
func RangeBasic() {
	pows := [4]int{0, 2, 4, 8}

	for i, v := range pows {
		fmt.Printf("%v -> %v, ", i, v)
	}
	fmt.Println()
}

//Range1 example
func Range1() {
	pows := [4]int{0, 2, 4, 8}
	for _, v := range pows {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
