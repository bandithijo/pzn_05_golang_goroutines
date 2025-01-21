package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)	
		channel <- "Rizqi Nur Assyaufi"	
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizqi Nur Assyaufi"
	fmt.Println("Selesai Mengirim Data ke Channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizqi Nur Assyaufi"
	fmt.Println("Selesai Mengirim Data ke Channel")
}

func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
	fmt.Println("Selesai Menerima Data dari Channel")
}

func TestInOutChanel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Rizqi"
		channel <- "Nur"
		channel <- "Assyaufi"
	}()

	go func() {
		for i := 0; i < cap(channel); i++ {
			fmt.Println(<- channel)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		fmt.Println("Selesai Mengirim Data ke Channel")
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("Selesai Menerima Data dari Channel")
}
