package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, num  int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello", num)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	totalHello := 0
	for i := 0; i < 100; i++ {
		go RunAsynchronous(group, i)
		totalHello++;
	}

	group.Wait()
	fmt.Println("Complete\nTotal Hello:", totalHello)
}
