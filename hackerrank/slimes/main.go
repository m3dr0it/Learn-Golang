package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var winner = map[int]int{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	test := strings.Split(text, "")
	fmt.Println(test)
	left := 2

	q, w := winner[left]
	fmt.Print("te")
	fmt.Println(q)

	fmt.Println(w)

	fmt.Println(left)

	winner[1] = 10
	winner[1] = winner[1] + 10

	fmt.Println(winner[1])

}
