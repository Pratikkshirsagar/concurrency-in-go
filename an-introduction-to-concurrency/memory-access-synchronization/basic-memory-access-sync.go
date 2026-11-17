package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int

	go func() {
		mu.Lock()
		counter++
		mu.Unlock()
	}()

	mu.Lock()
	if counter == 0 {
		fmt.Printf("The value is %v\n", counter)
	} else {
		fmt.Printf("The value is %v\n", counter)
	}
	mu.Unlock()
}



