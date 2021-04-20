package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go hello()
	fmt.Println("main goroutine done!")
	wg.Wait()
}

func hello() {
	fmt.Println("hello Goroutine!")
	wg.Done()
}
