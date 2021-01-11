package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//Go1 test
func Go1() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

//Channel1 test
func Channel1() {
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Printf("%T\n", s)
	c := make(chan int)
	go sum(s[:(len(s)/2)-1], c)
	go sum(s[len(s)/2:], c)

	fmt.Println(len(c))
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

//Channel2 test
func Channel2() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3
	fmt.Printf("len: %v cap: %v\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("F -> ")
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//Channel3 test
func Channel3() {
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)

	for i := 0; i < 10; i++ {
		i, closed := <-c
		fmt.Println("R ->", len(c), i, closed)
	}

}

//Select1  test
func Select1() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

//DefaultSelect test
func DefaultSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case x := <-tick:
			fmt.Println("tick.", x)
		case x := <-boom:
			fmt.Println("BOOM!", x)
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

//Concurrency example
func Concurrency() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
