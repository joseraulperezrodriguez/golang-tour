package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var mutex sync.Mutex

func concurrentIncrease(val *int, inc int, ops chan int) {
	mutex.Lock()
	(*val) += inc
	ops <- (*val)
	mutex.Unlock()
}

func ConcurrentWrite() {
	inc := []int{2, 5, 7}
	val := 0
	ops := make(chan int, 100)
	iterations := 100
	for i := 0; i < iterations; i++ {
		go concurrentIncrease(&val, inc[i%3], ops)
	}
	for i := 0; i < iterations; i++ {
		fmt.Printf("%v ", <-ops)
	}
}

func WaitGroup() {
	var wg sync.WaitGroup

	function := (func() {
		fmt.Print("1 ")
	})

	wg.Go(function)
	wg.Go(function)
	wg.Wait()
	fmt.Println("Wait Group Done")
}

func Write(val *int, inc int, lock *sync.RWMutex) {
	lock.Lock()
	(*val) += inc
	lock.Unlock()
}

func Read(val *int, lock *sync.RWMutex) {
	lock.RLock()
	fmt.Printf("%d ", (*val))
	lock.RUnlock()
}

var rn = rand.New(rand.NewSource(time.Now().UnixNano()))

func IsWrite() bool {
	return (rn.Int() % 2) == 0
}

func ReadWriteMutex() {
	var rwMutex sync.RWMutex
	val := 0
	inc := []int{1, 2, 3}
	for i := 0; i < 100; i++ {
		if IsWrite() {
			Write(&val, inc[i%3], &rwMutex)
		} else {
			Read(&val, &rwMutex)
		}
	}
}

func OnceUsage() {
	var once sync.Once

	for i := 0; i < 5; i++ {
		once.Do(func() {
			fmt.Println("This will only be printed once.")
		})
	}
}

func PoolUsage() {
	var pool sync.Pool

	pool.New = func() any {
		return make([]int, 0)
	}
	for i := 0; i < 5; i++ {
		val := pool.Get()
		array := val.([]int)
		array = append(array, i)
		fmt.Println(array[0])
		pool.Put(val)
	}
}

func AtomicOperations() {
	var c1 int64 = 10
	atomic.StoreInt64(&c1, 5)
	v := atomic.LoadInt64(&c1)
	fmt.Printf("Counter value: %d\n", v)

	var counter atomic.Int64
	counter.Add(1)
	fmt.Printf("Counter value: %d\n", counter.Load())
}
