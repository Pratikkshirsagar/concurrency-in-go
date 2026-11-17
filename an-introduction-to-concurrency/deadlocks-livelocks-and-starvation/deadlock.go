package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

// 1. Mutual exclusion: only one holder at a time, and the others have to wait.
// 2. Hold and wait: the process keeps what it already has while it waits for more.
// 3. No preemption: nobody can take a resource away. Only the holder can release it.
// 4. Circular wait: the waits form a closed loop, like A → B → A.
func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
