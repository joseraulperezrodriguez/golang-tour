package goroutines

import "fmt"

//MainGoRoutines test
func MainGoRoutines() {
	fmt.Println("Start MainGoRoutines >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	Go1()
	fmt.Println("-----------------------------------------------------")
	Channel1()
	fmt.Println("-----------------------------------------------------")
	Channel2()
	fmt.Println("-----------------------------------------------------")
	Channel3()
	fmt.Println("-----------------------------------------------------")
	Select1()
	fmt.Println("-----------------------------------------------------")
	DefaultSelect()
	fmt.Println("-----------------------------------------------------")
	Concurrency()
	fmt.Println("End MainGoRoutines <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println()

}
