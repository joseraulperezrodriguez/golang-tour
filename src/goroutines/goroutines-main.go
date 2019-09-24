package goroutines

import "fmt"

//MainGoRoutines test
func MainGoRoutines() {
	fmt.Println("Start MainGoRoutines >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	Go1()
	Channel1()
	Channel2()
	Channel3()
	Select1()
	DefaultSelect()
	Concurrency()
	fmt.Println("End MainGoRoutines <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println()

}
