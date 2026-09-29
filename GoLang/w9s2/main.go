package main

import (
	"fmt"

	"sync"
)

var count = 0

var wg sync.WaitGroup

func increment(num int) {
	defer wg.Done()
	count++
	fmt.Printf("Incremented: %d\n", num)
}

func main() {
	for i := 0; i < 10000; i++ {
		go increment(i)
		wg.Add(1)
	}
	wg.Wait()

}
