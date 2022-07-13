package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrintWg(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func printIt(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)
		wg.Add(1)
		go doPrintWg(&wg, data)
	}

	wg.Wait()
}
