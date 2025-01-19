package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello, Go!")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("Ups!")

	time.Sleep(1 * time.Second)
}
