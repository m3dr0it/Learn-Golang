package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chan1 := make(chan int, 100)

	for port := 2000; port < 4000; port++ {
		wg.Add(1)
		chan1 <- port
		go worker(chan1, &wg)
	}
	wg.Wait()
}

func worker(ports chan int, wg *sync.WaitGroup) {
	var a uint8
	a = 0
	log.Println(a)

	for port := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", port)
		_, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Printf("port %d closed \n", port)
		} else {
			fmt.Printf("port %d opened \n", port)
		}
		wg.Done()
	}
}
