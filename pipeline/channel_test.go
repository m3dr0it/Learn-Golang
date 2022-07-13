package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	fmt.Println("start")
	chan1 := make(chan string)
	close(chan1)
	test, isClose := <-chan1
	fmt.Println(test)
	fmt.Println(isClose)
	// chan1 <- "1"
	// chan1 <- "2"

	go func() {
		for elem := range chan1 {
			fmt.Println(elem)
		}
	}()

	time.Sleep(1 * time.Second)
}
