package errors

import (
	"fmt"
	"runtime/debug"
	"time"
)

//MyError custom error
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"Some explotion",
	}
}

//Error1 test
func Error1() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

//Error2 test
func Error2() {
	fmt.Println("Current stack =========================================================")
	debug.PrintStack()
}

func panicking() {
	panic("I enter in panic")
}

//Panic1  test
func Panic1() {
	defer func() {
		recovered := recover()
		fmt.Printf("Try to handle panic %v and error %T\n", recovered, recovered)
	}()
	panicking()
}
