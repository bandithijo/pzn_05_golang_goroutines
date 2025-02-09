package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var group sync.WaitGroup
	var counter int64 = 0

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)

			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
