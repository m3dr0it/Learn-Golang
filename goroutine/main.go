package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "testing")

	print(4, "hehe")

	var input string
	fmt.Scanln(&input)

	fmt.Println(input)
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(message)
	}
}
