package concurrency

import "fmt"

//MainConcurrency test

func MainConcurrency() {
	fmt.Println("Start MainConcurrency >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	ConcurrentWrite()
	fmt.Println("\n-------------------------")
	WaitGroup()
	fmt.Println("\n-------------------------")
	ReadWriteMutex()
	fmt.Println("\n-------------------------")
	OnceUsage()
	fmt.Println("\n-------------------------")
	PoolUsage()
	fmt.Println("\n-------------------------")
	AtomicOperations()
	fmt.Println("End MainConcurrency <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println()

}
