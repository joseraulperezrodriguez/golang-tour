package flowcontrol

import (
	"fmt"
	"runtime"
	"time"
)

//ClassicFor just use the standard way to use for
func ClassicFor() {
	sum := 0
	for i := 1; i < 5; i++ {
		fmt.Printf("%v ", i)
		sum += i
	}
	fmt.Printf(" => sum: %v\n", sum)
}

//WhileFor while style for
func WhileFor() {
	sum := 10
	for sum >= 0 {
		sum--
	}
	fmt.Println(sum)
}

//If test
func If() {
	if true {
		fmt.Println("If")
	}
}

//IfVar just test the variable assignmente in Go
func IfVar() {
	if x := 1; x < 5 {
		fmt.Printf("%v < 5\n", x)
	} else {
		fmt.Printf("%v >= 5\n", x)
	}
}

//SwitchBasic for learning to use SwitchBasic
func SwitchBasic() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("darwin")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

//Switch1 test
func Switch1() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//SwitchNoCondition with no condition
func SwitchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

//Defer first use
func Defer() {
	defer fmt.Println("Execute Defer Code")
	fmt.Println("End Function")
}

//DeferStack shows the defer stack
func DeferStack() {
	defer fmt.Println()
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v", i)
	}
	fmt.Print("Reverse Numbers: ")
}
