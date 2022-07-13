package main

import "fmt"

func main() {
	var testChann = make(chan int)

	var printIt = func(a int) {
		testChann <- a
	}

	go printIt(1)
	go printIt(2)
	go printIt(3)
	go printIt(4)

	var receiver = <-testChann
	fmt.Println(receiver)

	var receiver1 = <-testChann
	fmt.Println(receiver1)

	var receiver2 = <-testChann
	fmt.Println(receiver2)

	var receiver3 = <-testChann
	fmt.Println(receiver3)
}
